package repositories

import (
	"context"
	"database/sql"
	"time"

	"obp-api-backend/internal/models"
	"obp-api-backend/pkg/db"
)

type metricsRepository struct {
	db *sql.DB
}

func NewMetricsRepository() MetricsRepository {
	return &metricsRepository{
		db: db.GetDB(),
	}
}

func (r *metricsRepository) RecordAPICall(ctx context.Context, consumerID, userID, url string, duration int64) error {
	query := `INSERT INTO api_metrics (consumer_id, user_id, url, duration) VALUES (?, ?, ?, ?)`
	
	_, err := r.db.ExecContext(ctx, query, consumerID, userID, url, duration)
	return err
}

func (r *metricsRepository) GetMetrics(ctx context.Context, fromDate, toDate time.Time, consumerID, userID, url string) (*models.APIMetrics, error) {
	query := `SELECT 
				COUNT(*) as total_calls,
				AVG(duration) as average_duration,
				MIN(duration) as min_duration,
				MAX(duration) as max_duration,
				SUM(duration) as total_duration
			  FROM api_metrics 
			  WHERE date BETWEEN ? AND ?`
	
	args := []interface{}{fromDate.Format(time.RFC3339), toDate.Format(time.RFC3339)}
	
	if consumerID != "" {
		query += " AND consumer_id = ?"
		args = append(args, consumerID)
	}
	if userID != "" {
		query += " AND user_id = ?"
		args = append(args, userID)
	}
	if url != "" {
		query += " AND url LIKE ?"
		args = append(args, "%"+url+"%")
	}
	
	metrics := &models.APIMetrics{}
	err := r.db.QueryRowContext(ctx, query, args...).Scan(
		&metrics.TotalCalls,
		&metrics.AverageDuration,
		&metrics.MinDuration,
		&metrics.MaxDuration,
		&metrics.TotalDuration,
	)
	
	return metrics, err
}

func (r *metricsRepository) GetTopAPIs(ctx context.Context, fromDate, toDate time.Time, limit int) ([]*models.APIRanking, error) {
	query := `SELECT 
				url,
				COUNT(*) as call_count,
				AVG(duration) as avg_duration,
				SUM(duration) as total_duration
			  FROM api_metrics 
			  WHERE date BETWEEN ? AND ?
			  GROUP BY url
			  ORDER BY call_count DESC, avg_duration ASC
			  LIMIT ?`
	
	rows, err := r.db.QueryContext(ctx, query, fromDate.Format(time.RFC3339), toDate.Format(time.RFC3339), limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var rankings []*models.APIRanking
	rank := 1
	for rows.Next() {
		ranking := &models.APIRanking{}
		
		err := rows.Scan(
			&ranking.URL,
			&ranking.CallCount,
			&ranking.AverageDuration,
			&ranking.TotalDuration,
		)
		if err != nil {
			return nil, err
		}
		
		ranking.Rank = rank
		rankings = append(rankings, ranking)
		rank++
	}
	
	return rankings, rows.Err()
}

func (r *metricsRepository) GetTopConsumers(ctx context.Context, fromDate, toDate time.Time, limit int) ([]*models.ConsumerRanking, error) {
	query := `SELECT 
				consumer_id,
				COUNT(*) as total_calls,
				COUNT(DISTINCT url) as unique_apis_used,
				AVG(duration) as avg_response_time,
				SUM(duration) as total_duration
			  FROM api_metrics 
			  WHERE date BETWEEN ? AND ? AND consumer_id IS NOT NULL
			  GROUP BY consumer_id
			  ORDER BY total_calls DESC, unique_apis_used DESC
			  LIMIT ?`
	
	rows, err := r.db.QueryContext(ctx, query, fromDate.Format(time.RFC3339), toDate.Format(time.RFC3339), limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var rankings []*models.ConsumerRanking
	rank := 1
	for rows.Next() {
		ranking := &models.ConsumerRanking{}
		
		err := rows.Scan(
			&ranking.ConsumerID,
			&ranking.TotalCalls,
			&ranking.UniqueAPIs,
			&ranking.AverageResponseTime,
			&ranking.TotalDuration,
		)
		if err != nil {
			return nil, err
		}
		
		ranking.Rank = rank
		rankings = append(rankings, ranking)
		rank++
	}
	
	return rankings, rows.Err()
}
