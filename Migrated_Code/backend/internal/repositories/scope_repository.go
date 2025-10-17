package repositories

import (
	"context"
	"obp-api-backend/internal/models"
	"obp-api-backend/pkg/db"
)

type scopeRepository struct{}

func NewScopeRepository() ScopeRepository {
	return &scopeRepository{}
}

func (r *scopeRepository) Create(ctx context.Context, scope *models.Scope) error {
	database := db.GetDB()
	query := `
		INSERT INTO scopes (scope_id, consumer_id, role_name, bank_id, created_at)
		VALUES (?, ?, ?, ?, ?)
	`
	_, err := database.ExecContext(ctx, query,
		scope.ScopeID,
		scope.ConsumerID,
		scope.RoleName,
		scope.BankID,
		scope.CreatedAt,
	)
	return err
}

func (r *scopeRepository) GetByConsumerID(ctx context.Context, consumerID string) ([]*models.Scope, error) {
	database := db.GetDB()
	query := `
		SELECT scope_id, consumer_id, role_name, bank_id, created_at
		FROM scopes
		WHERE consumer_id = ?
	`

	rows, err := database.QueryContext(ctx, query, consumerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var scopes []*models.Scope
	for rows.Next() {
		scope := &models.Scope{}
		err := rows.Scan(
			&scope.ScopeID,
			&scope.ConsumerID,
			&scope.RoleName,
			&scope.BankID,
			&scope.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		scopes = append(scopes, scope)
	}

	return scopes, rows.Err()
}

func (r *scopeRepository) HasRole(ctx context.Context, consumerID, roleName, bankID string) (bool, error) {
	database := db.GetDB()

	query := `
		SELECT COUNT(*) FROM scopes
		WHERE consumer_id = ? AND role_name = ? AND (bank_id = ? OR bank_id = '')
	`

	var count int
	err := database.QueryRowContext(ctx, query, consumerID, roleName, bankID).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *scopeRepository) Delete(ctx context.Context, scopeID string) error {
	database := db.GetDB()
	query := `DELETE FROM scopes WHERE scope_id = ?`
	_, err := database.ExecContext(ctx, query, scopeID)
	return err
}
