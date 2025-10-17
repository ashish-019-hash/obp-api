package services

import (
	"context"
	"fmt"
	"time"

	"obp-api-backend/internal/repositories"
)

type RateLimitingService interface {
	CheckRateLimit(ctx context.Context, consumerKey string, currentTime time.Time) (*RateLimitResult, error)
	IncrementCounter(ctx context.Context, consumerKey string, currentTime time.Time) error
	GetRemainingCalls(ctx context.Context, consumerKey string, period string, currentTime time.Time) (int, error)
}

type RateLimitResult struct {
	IsAllowed      bool
	RemainingCalls map[string]int
	ResetTimes     map[string]time.Time
	ExceededPeriod string
}

type rateLimitingService struct {
	rateLimitRepo repositories.RateLimitRepository
	limits        map[string]int
}

func NewRateLimitingService(rateLimitRepo repositories.RateLimitRepository) RateLimitingService {
	return &rateLimitingService{
		rateLimitRepo: rateLimitRepo,
		limits: map[string]int{
			"second": 10,
			"minute": 100,
			"hour":   1000,
			"day":    10000,
			"week":   50000,
			"month":  200000,
			"year":   1000000,
		},
	}
}

func (s *rateLimitingService) CheckRateLimit(ctx context.Context, consumerKey string, currentTime time.Time) (*RateLimitResult, error) {
	result := &RateLimitResult{
		IsAllowed:      true,
		RemainingCalls: make(map[string]int),
		ResetTimes:     make(map[string]time.Time),
	}

	periods := []string{"second", "minute", "hour", "day", "week", "month", "year"}

	for _, period := range periods {
		periodKey := s.generatePeriodKey(consumerKey, period, currentTime)
		limit := s.limits[period]

		currentCount, err := s.rateLimitRepo.GetCounter(ctx, consumerKey, period, periodKey)
		if err != nil {
			currentCount = 0
		}

		if currentCount >= limit {
			result.IsAllowed = false
			result.ExceededPeriod = period
			result.RemainingCalls[period] = 0
		} else {
			result.RemainingCalls[period] = limit - currentCount
		}

		result.ResetTimes[period] = s.getResetTime(period, currentTime)
	}

	return result, nil
}

func (s *rateLimitingService) IncrementCounter(ctx context.Context, consumerKey string, currentTime time.Time) error {
	periods := []string{"second", "minute", "hour", "day", "week", "month", "year"}

	for _, period := range periods {
		periodKey := s.generatePeriodKey(consumerKey, period, currentTime)
		limit := s.limits[period]

		_, err := s.rateLimitRepo.IncrementCounter(ctx, consumerKey, period, periodKey, limit)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *rateLimitingService) GetRemainingCalls(ctx context.Context, consumerKey string, period string, currentTime time.Time) (int, error) {
	periodKey := s.generatePeriodKey(consumerKey, period, currentTime)
	limit := s.limits[period]

	currentCount, err := s.rateLimitRepo.GetCounter(ctx, consumerKey, period, periodKey)
	if err != nil {
		return limit, nil
	}

	remaining := limit - currentCount
	if remaining < 0 {
		remaining = 0
	}

	return remaining, nil
}

func (s *rateLimitingService) generatePeriodKey(consumerKey, period string, currentTime time.Time) string {
	switch period {
	case "second":
		return fmt.Sprintf("%s:%s:%d", consumerKey, period, currentTime.Unix())
	case "minute":
		return fmt.Sprintf("%s:%s:%d", consumerKey, period, currentTime.Unix()/60)
	case "hour":
		return fmt.Sprintf("%s:%s:%d", consumerKey, period, currentTime.Unix()/3600)
	case "day":
		return fmt.Sprintf("%s:%s:%s", consumerKey, period, currentTime.Format("2006-01-02"))
	case "week":
		year, week := currentTime.ISOWeek()
		return fmt.Sprintf("%s:%s:%d-%d", consumerKey, period, year, week)
	case "month":
		return fmt.Sprintf("%s:%s:%s", consumerKey, period, currentTime.Format("2006-01"))
	case "year":
		return fmt.Sprintf("%s:%s:%d", consumerKey, period, currentTime.Year())
	default:
		return fmt.Sprintf("%s:%s:%d", consumerKey, period, currentTime.Unix())
	}
}

func (s *rateLimitingService) getResetTime(period string, currentTime time.Time) time.Time {
	switch period {
	case "second":
		return currentTime.Truncate(time.Second).Add(time.Second)
	case "minute":
		return currentTime.Truncate(time.Minute).Add(time.Minute)
	case "hour":
		return currentTime.Truncate(time.Hour).Add(time.Hour)
	case "day":
		return time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day()+1, 0, 0, 0, 0, currentTime.Location())
	case "week":
		weekday := currentTime.Weekday()
		daysUntilMonday := (7 - int(weekday) + int(time.Monday)) % 7
		if daysUntilMonday == 0 {
			daysUntilMonday = 7
		}
		return currentTime.AddDate(0, 0, daysUntilMonday).Truncate(24 * time.Hour)
	case "month":
		return time.Date(currentTime.Year(), currentTime.Month()+1, 1, 0, 0, 0, 0, currentTime.Location())
	case "year":
		return time.Date(currentTime.Year()+1, 1, 1, 0, 0, 0, 0, currentTime.Location())
	default:
		return currentTime.Add(time.Hour)
	}
}
