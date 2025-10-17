package repositories

import (
	"context"
	"database/sql"

	"obp-api-backend/internal/models"
	"obp-api-backend/pkg/db"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository() UserRepository {
	return &userRepository{
		db: db.GetDB(),
	}
}

func (r *userRepository) Create(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users (user_id, provider, provider_id, username, email) VALUES (?, ?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query,
		user.UserID,
		user.Provider,
		"",
		user.Name,
		user.Email,
	)
	return err
}

func (r *userRepository) GetByID(ctx context.Context, userID string) (*models.User, error) {
	query := `SELECT user_id, provider, provider_id, username, email FROM users WHERE user_id = ?`

	user := &models.User{}
	var providerId string

	err := r.db.QueryRowContext(ctx, query, userID).Scan(
		&user.UserID,
		&user.Provider,
		&providerId,
		&user.Name,
		&user.Email,
	)

	return user, err
}

func (r *userRepository) GetByProvider(ctx context.Context, provider, providerID string) (*models.User, error) {
	query := `SELECT user_id, provider, provider_id, username, email FROM users WHERE provider = ? AND provider_id = ?`

	user := &models.User{}
	var providerIdFromDB string

	err := r.db.QueryRowContext(ctx, query, provider, providerID).Scan(
		&user.UserID,
		&user.Provider,
		&providerIdFromDB,
		&user.Name,
		&user.Email,
	)

	return user, err
}

func (r *userRepository) Update(ctx context.Context, user *models.User) error {
	query := `UPDATE users SET provider = ?, username = ?, email = ? WHERE user_id = ?`

	_, err := r.db.ExecContext(ctx, query,
		user.Provider,
		user.Name,
		user.Email,
		user.UserID,
	)
	return err
}

func (r *userRepository) Delete(ctx context.Context, userID string) error {
	query := `DELETE FROM users WHERE user_id = ?`
	_, err := r.db.ExecContext(ctx, query, userID)
	return err
}
