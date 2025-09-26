package services

import (
	"sync"
	"time"

	"github.com/ashish-019-hash/obp-api-backend/internal/models"
	"gorm.io/gorm"
)

type RateLimiter struct {
	windows       map[string]*SlidingWindow
	mutex         sync.RWMutex
	configService *ConfigService
	db            *gorm.DB
}

type SlidingWindow struct {
	requests []time.Time
	mutex    sync.Mutex
}

func NewRateLimiter(configService *ConfigService, db *gorm.DB) *RateLimiter {
	return &RateLimiter{
		windows:       make(map[string]*SlidingWindow),
		configService: configService,
		db:            db,
	}
}

func (rl *RateLimiter) IsLimited(identifier string) bool {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()
	
	window, exists := rl.windows[identifier]
	if !exists {
		window = &SlidingWindow{requests: make([]time.Time, 0)}
		rl.windows[identifier] = window
	}
	
	perMinute, perHour := rl.getRateLimits(identifier)
	return window.checkLimit(perMinute, perHour)
}

func (rl *RateLimiter) getRateLimits(identifier string) (int, int) {
	var consumer models.Consumer
	if err := rl.db.Where("consumer_id = ?", identifier).First(&consumer).Error; err == nil {
		if rateLimit, err := rl.configService.GetConsumerRateLimit(identifier); err == nil {
			return rateLimit.RequestsPerMinute, rateLimit.RequestsPerHour
		}
	}
	
	perMinute := rl.configService.GetConfigInt("rate.limiting.anonymous.per.minute", 100)
	perHour := rl.configService.GetConfigInt("rate.limiting.anonymous.per.hour", 1000)
	
	return perMinute, perHour
}

func (rl *RateLimiter) SetConsumerLimits(consumerID string, perMinute, perHour, perDay int) error {
	return rl.configService.SetConsumerRateLimit(consumerID, perMinute, perHour, perDay)
}

func (rl *RateLimiter) GetLimits(identifier string) (int, int) {
	return rl.getRateLimits(identifier)
}

func (sw *SlidingWindow) checkLimit(perMinute, perHour int) bool {
	sw.mutex.Lock()
	defer sw.mutex.Unlock()
	
	now := time.Now()
	oneMinuteAgo := now.Add(-time.Minute)
	oneHourAgo := now.Add(-time.Hour)
	
	var validRequests []time.Time
	minuteCount := 0
	hourCount := 0
	
	for _, req := range sw.requests {
		if req.After(oneHourAgo) {
			validRequests = append(validRequests, req)
			hourCount++
			if req.After(oneMinuteAgo) {
				minuteCount++
			}
		}
	}
	
	sw.requests = validRequests
	
	if minuteCount >= perMinute || hourCount >= perHour {
		return true
	}
	
	sw.requests = append(sw.requests, now)
	return false
}

func (rl *RateLimiter) Reset(identifier string) {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()
	
	delete(rl.windows, identifier)
}

func (rl *RateLimiter) GetStats(identifier string) (int, int) {
	rl.mutex.RLock()
	defer rl.mutex.RUnlock()
	
	window, exists := rl.windows[identifier]
	if !exists {
		return 0, 0
	}
	
	window.mutex.Lock()
	defer window.mutex.Unlock()
	
	now := time.Now()
	oneMinuteAgo := now.Add(-time.Minute)
	oneHourAgo := now.Add(-time.Hour)
	
	minuteCount := 0
	hourCount := 0
	
	for _, req := range window.requests {
		if req.After(oneHourAgo) {
			hourCount++
			if req.After(oneMinuteAgo) {
				minuteCount++
			}
		}
	}
	
	return minuteCount, hourCount
}
