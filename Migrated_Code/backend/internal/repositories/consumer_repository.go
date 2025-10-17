package repositories

import (
	"context"
	"database/sql"
	"obp-api-backend/internal/models"
	"obp-api-backend/pkg/db"
)

type consumerRepository struct{}

func NewConsumerRepository() ConsumerRepository {
	return &consumerRepository{}
}

func (r *consumerRepository) Create(ctx context.Context, consumer *models.Consumer) error {
	database := db.GetDB()
	query := `
		INSERT INTO consumers (
			consumer_id, consumer_key, consumer_secret, is_active, name, app_type, description,
			developer_email, redirect_url, created_by_user_id, client_certificate, company, logo_url,
			per_second_call_limit, per_minute_call_limit, per_hour_call_limit, per_day_call_limit,
			per_week_call_limit, per_month_call_limit, created_at, updated_at
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := database.ExecContext(ctx, query,
		consumer.ConsumerID,
		consumer.ConsumerKey,
		consumer.ConsumerSecret,
		consumer.IsActive,
		consumer.Name,
		consumer.AppType,
		consumer.Description,
		consumer.DeveloperEmail,
		consumer.RedirectURL,
		consumer.CreatedByUserID,
		consumer.ClientCertificate,
		consumer.Company,
		consumer.LogoURL,
		consumer.PerSecondCallLimit,
		consumer.PerMinuteCallLimit,
		consumer.PerHourCallLimit,
		consumer.PerDayCallLimit,
		consumer.PerWeekCallLimit,
		consumer.PerMonthCallLimit,
		consumer.CreatedAt,
		consumer.UpdatedAt,
	)
	return err
}

func (r *consumerRepository) GetByConsumerKey(ctx context.Context, consumerKey string) (*models.Consumer, error) {
	database := db.GetDB()
	query := `
		SELECT consumer_id, consumer_key, consumer_secret, is_active, name, app_type, description,
		       developer_email, redirect_url, created_by_user_id, client_certificate, company, logo_url,
		       per_second_call_limit, per_minute_call_limit, per_hour_call_limit, per_day_call_limit,
		       per_week_call_limit, per_month_call_limit, created_at, updated_at
		FROM consumers
		WHERE consumer_key = ?
	`

	consumer := &models.Consumer{}
	err := database.QueryRowContext(ctx, query, consumerKey).Scan(
		&consumer.ConsumerID,
		&consumer.ConsumerKey,
		&consumer.ConsumerSecret,
		&consumer.IsActive,
		&consumer.Name,
		&consumer.AppType,
		&consumer.Description,
		&consumer.DeveloperEmail,
		&consumer.RedirectURL,
		&consumer.CreatedByUserID,
		&consumer.ClientCertificate,
		&consumer.Company,
		&consumer.LogoURL,
		&consumer.PerSecondCallLimit,
		&consumer.PerMinuteCallLimit,
		&consumer.PerHourCallLimit,
		&consumer.PerDayCallLimit,
		&consumer.PerWeekCallLimit,
		&consumer.PerMonthCallLimit,
		&consumer.CreatedAt,
		&consumer.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return consumer, err
}

func (r *consumerRepository) GetByID(ctx context.Context, consumerID string) (*models.Consumer, error) {
	database := db.GetDB()
	query := `
		SELECT consumer_id, consumer_key, consumer_secret, is_active, name, app_type, description,
		       developer_email, redirect_url, created_by_user_id, client_certificate, company, logo_url,
		       per_second_call_limit, per_minute_call_limit, per_hour_call_limit, per_day_call_limit,
		       per_week_call_limit, per_month_call_limit, created_at, updated_at
		FROM consumers
		WHERE consumer_id = ?
	`

	consumer := &models.Consumer{}
	err := database.QueryRowContext(ctx, query, consumerID).Scan(
		&consumer.ConsumerID,
		&consumer.ConsumerKey,
		&consumer.ConsumerSecret,
		&consumer.IsActive,
		&consumer.Name,
		&consumer.AppType,
		&consumer.Description,
		&consumer.DeveloperEmail,
		&consumer.RedirectURL,
		&consumer.CreatedByUserID,
		&consumer.ClientCertificate,
		&consumer.Company,
		&consumer.LogoURL,
		&consumer.PerSecondCallLimit,
		&consumer.PerMinuteCallLimit,
		&consumer.PerHourCallLimit,
		&consumer.PerDayCallLimit,
		&consumer.PerWeekCallLimit,
		&consumer.PerMonthCallLimit,
		&consumer.CreatedAt,
		&consumer.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return consumer, err
}

func (r *consumerRepository) Update(ctx context.Context, consumer *models.Consumer) error {
	database := db.GetDB()
	query := `
		UPDATE consumers
		SET consumer_key = ?, consumer_secret = ?, is_active = ?, name = ?, app_type = ?, 
		    description = ?, developer_email = ?, redirect_url = ?, created_by_user_id = ?,
		    client_certificate = ?, company = ?, logo_url = ?, per_second_call_limit = ?,
		    per_minute_call_limit = ?, per_hour_call_limit = ?, per_day_call_limit = ?,
		    per_week_call_limit = ?, per_month_call_limit = ?, updated_at = ?
		WHERE consumer_id = ?
	`
	_, err := database.ExecContext(ctx, query,
		consumer.ConsumerKey,
		consumer.ConsumerSecret,
		consumer.IsActive,
		consumer.Name,
		consumer.AppType,
		consumer.Description,
		consumer.DeveloperEmail,
		consumer.RedirectURL,
		consumer.CreatedByUserID,
		consumer.ClientCertificate,
		consumer.Company,
		consumer.LogoURL,
		consumer.PerSecondCallLimit,
		consumer.PerMinuteCallLimit,
		consumer.PerHourCallLimit,
		consumer.PerDayCallLimit,
		consumer.PerWeekCallLimit,
		consumer.PerMonthCallLimit,
		consumer.UpdatedAt,
		consumer.ConsumerID,
	)
	return err
}

func (r *consumerRepository) Delete(ctx context.Context, consumerID string) error {
	database := db.GetDB()
	query := `UPDATE consumers SET is_active = FALSE WHERE consumer_id = ?`
	_, err := database.ExecContext(ctx, query, consumerID)
	return err
}
