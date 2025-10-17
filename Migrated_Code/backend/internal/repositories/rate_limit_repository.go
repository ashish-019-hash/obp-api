package repositories

import (
	"context"
	"database/sql"
	"time"

	"obp-api-backend/pkg/db"
)

type rateLimitRepository struct {
	db *sql.DB
}

func NewRateLimitRepository() RateLimitRepository {
	return &rateLimitRepository{
		db: db.GetDB(),
	}
}

func (r *rateLimitRepository) IncrementCounter(ctx context.Context, consumerKey, period, periodKey string, limit int) (int, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	var currentCount int
	query := `SELECT call_count FROM rate_limits WHERE consumer_key = ? AND period = ? AND period_key = ?`
	err = tx.QueryRowContext(ctx, query, consumerKey, period, periodKey).Scan(&currentCount)

	if err == sql.ErrNoRows {
		insertQuery := `INSERT INTO rate_limits (consumer_key, period, period_key, call_count, limit_value, reset_time) 
						VALUES (?, ?, ?, 1, ?, ?)`
		resetTime := r.calculateResetTime(period)
		_, err = tx.ExecContext(ctx, insertQuery, consumerKey, period, periodKey, limit, resetTime)
		if err != nil {
			return 0, err
		}
		currentCount = 1
	} else if err != nil {
		return 0, err
	} else {
		updateQuery := `UPDATE rate_limits SET call_count = call_count + 1 WHERE consumer_key = ? AND period = ? AND period_key = ?`
		_, err = tx.ExecContext(ctx, updateQuery, consumerKey, period, periodKey)
		if err != nil {
			return 0, err
		}
		currentCount++
	}

	return currentCount, tx.Commit()
}

func (r *rateLimitRepository) GetCounter(ctx context.Context, consumerKey, period, periodKey string) (int, error) {
	query := `SELECT call_count FROM rate_limits WHERE consumer_key = ? AND period = ? AND period_key = ?`

	var count int
	err := r.db.QueryRowContext(ctx, query, consumerKey, period, periodKey).Scan(&count)
	if err == sql.ErrNoRows {
		return 0, nil
	}

	return count, err
}

func (r *rateLimitRepository) ResetCounter(ctx context.Context, consumerKey, period, periodKey string) error {
	query := `UPDATE rate_limits SET call_count = 0 WHERE consumer_key = ? AND period = ? AND period_key = ?`
	_, err := r.db.ExecContext(ctx, query, consumerKey, period, periodKey)
	return err
}

func (r *rateLimitRepository) calculateResetTime(period string) time.Time {
	now := time.Now()

	switch period {
	case "second":
		return now.Truncate(time.Second).Add(time.Second)
	case "minute":
		return now.Truncate(time.Minute).Add(time.Minute)
	case "hour":
		return now.Truncate(time.Hour).Add(time.Hour)
	case "day":
		return time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
	case "week":
		weekday := now.Weekday()
		daysUntilMonday := (7 - int(weekday) + int(time.Monday)) % 7
		if daysUntilMonday == 0 {
			daysUntilMonday = 7
		}
		return now.AddDate(0, 0, daysUntilMonday).Truncate(24 * time.Hour)
	case "month":
		return time.Date(now.Year(), now.Month()+1, 1, 0, 0, 0, 0, now.Location())
	case "year":
		return time.Date(now.Year()+1, 1, 1, 0, 0, 0, 0, now.Location())
	default:
		return now.Add(time.Hour)
	}
}
