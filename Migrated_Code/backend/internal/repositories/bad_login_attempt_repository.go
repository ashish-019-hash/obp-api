package repositories

import (
	"context"
	"database/sql"
	"obp-api-backend/internal/models"
	"obp-api-backend/pkg/db"
	"time"
)

type badLoginAttemptRepository struct{}

func NewBadLoginAttemptRepository() BadLoginAttemptRepository {
	return &badLoginAttemptRepository{}
}

func (r *badLoginAttemptRepository) GetByUsernameProvider(ctx context.Context, username, provider string) (*models.BadLoginAttempt, error) {
	database := db.GetDB()
	query := `
		SELECT id, provider, username, last_failure_date, bad_attempts_since_last_success_or_reset
		FROM bad_login_attempts
		WHERE username = ? AND provider = ?
	`

	attempt := &models.BadLoginAttempt{}
	err := database.QueryRowContext(ctx, query, username, provider).Scan(
		&attempt.ID,
		&attempt.Provider,
		&attempt.Username,
		&attempt.LastFailureDate,
		&attempt.BadAttemptsSinceLastSuccessOrReset,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return attempt, err
}

func (r *badLoginAttemptRepository) IncrementAttempts(ctx context.Context, username, provider string) error {
	database := db.GetDB()
	query := `
		INSERT INTO bad_login_attempts (provider, username, last_failure_date, bad_attempts_since_last_success_or_reset)
		VALUES (?, ?, ?, 1)
		ON CONFLICT(provider, username) DO UPDATE SET
			last_failure_date = ?,
			bad_attempts_since_last_success_or_reset = bad_attempts_since_last_success_or_reset + 1
	`
	now := time.Now()
	_, err := database.ExecContext(ctx, query, provider, username, now, now)
	return err
}

func (r *badLoginAttemptRepository) ResetAttempts(ctx context.Context, username, provider string) error {
	database := db.GetDB()
	query := `
		UPDATE bad_login_attempts
		SET bad_attempts_since_last_success_or_reset = 0
		WHERE username = ? AND provider = ?
	`
	_, err := database.ExecContext(ctx, query, username, provider)
	return err
}
