package repositories

import (
	"context"
	"database/sql"
	"obp-api-backend/internal/models"
	"obp-api-backend/pkg/db"
)

type authUserRepository struct{}

func NewAuthUserRepository() AuthUserRepository {
	return &authUserRepository{}
}

func (r *authUserRepository) Create(ctx context.Context, user *models.AuthUser) error {
	database := db.GetDB()
	query := `
		INSERT INTO auth_users (user_id, username, email, password_hash, email_verified, validated, first_name, last_name, provider, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := database.ExecContext(ctx, query,
		user.UserID,
		user.Username,
		user.Email,
		user.PasswordHash,
		user.EmailVerified,
		user.Validated,
		user.FirstName,
		user.LastName,
		user.Provider,
		user.CreatedAt,
		user.UpdatedAt,
	)
	return err
}

func (r *authUserRepository) GetByUsername(ctx context.Context, username string) (*models.AuthUser, error) {
	database := db.GetDB()
	query := `
		SELECT user_id, username, email, password_hash, email_verified, validated, first_name, last_name, provider, created_at, updated_at
		FROM auth_users
		WHERE username = ?
	`

	user := &models.AuthUser{}
	err := database.QueryRowContext(ctx, query, username).Scan(
		&user.UserID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.EmailVerified,
		&user.Validated,
		&user.FirstName,
		&user.LastName,
		&user.Provider,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return user, err
}

func (r *authUserRepository) GetByEmail(ctx context.Context, email string) (*models.AuthUser, error) {
	database := db.GetDB()
	query := `
		SELECT user_id, username, email, password_hash, email_verified, validated, first_name, last_name, provider, created_at, updated_at
		FROM auth_users
		WHERE email = ?
	`

	user := &models.AuthUser{}
	err := database.QueryRowContext(ctx, query, email).Scan(
		&user.UserID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.EmailVerified,
		&user.Validated,
		&user.FirstName,
		&user.LastName,
		&user.Provider,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return user, err
}

func (r *authUserRepository) GetByID(ctx context.Context, userID string) (*models.AuthUser, error) {
	database := db.GetDB()
	query := `
		SELECT user_id, username, email, password_hash, email_verified, validated, first_name, last_name, provider, created_at, updated_at
		FROM auth_users
		WHERE user_id = ?
	`

	user := &models.AuthUser{}
	err := database.QueryRowContext(ctx, query, userID).Scan(
		&user.UserID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.EmailVerified,
		&user.Validated,
		&user.FirstName,
		&user.LastName,
		&user.Provider,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return user, err
}

func (r *authUserRepository) Update(ctx context.Context, user *models.AuthUser) error {
	database := db.GetDB()
	query := `
		UPDATE auth_users
		SET username = ?, email = ?, password_hash = ?, email_verified = ?, validated = ?, 
		    first_name = ?, last_name = ?, provider = ?, updated_at = ?
		WHERE user_id = ?
	`
	_, err := database.ExecContext(ctx, query,
		user.Username,
		user.Email,
		user.PasswordHash,
		user.EmailVerified,
		user.Validated,
		user.FirstName,
		user.LastName,
		user.Provider,
		user.UpdatedAt,
		user.UserID,
	)
	return err
}
