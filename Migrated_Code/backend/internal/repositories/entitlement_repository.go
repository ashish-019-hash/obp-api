package repositories

import (
	"context"
	"obp-api-backend/internal/models"
	"obp-api-backend/pkg/db"
)

type entitlementRepository struct{}

func NewEntitlementRepository() EntitlementRepository {
	return &entitlementRepository{}
}

func (r *entitlementRepository) Create(ctx context.Context, entitlement *models.Entitlement) error {
	database := db.GetDB()
	query := `
		INSERT INTO entitlements (entitlement_id, user_id, role_name, bank_id, created_by_user_id, created_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	_, err := database.ExecContext(ctx, query,
		entitlement.EntitlementID,
		entitlement.UserID,
		entitlement.RoleName,
		entitlement.BankID,
		entitlement.CreatedByUserID,
		entitlement.CreatedAt,
	)
	return err
}

func (r *entitlementRepository) GetByUserID(ctx context.Context, userID string) ([]*models.Entitlement, error) {
	database := db.GetDB()
	query := `
		SELECT entitlement_id, user_id, role_name, bank_id, created_by_user_id, created_at
		FROM entitlements
		WHERE user_id = ?
	`

	rows, err := database.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entitlements []*models.Entitlement
	for rows.Next() {
		ent := &models.Entitlement{}
		err := rows.Scan(
			&ent.EntitlementID,
			&ent.UserID,
			&ent.RoleName,
			&ent.BankID,
			&ent.CreatedByUserID,
			&ent.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		entitlements = append(entitlements, ent)
	}

	return entitlements, rows.Err()
}

func (r *entitlementRepository) HasRole(ctx context.Context, userID, roleName, bankID string) (bool, error) {
	database := db.GetDB()

	query := `
		SELECT COUNT(*) FROM entitlements
		WHERE user_id = ? AND role_name = ? AND (bank_id = ? OR bank_id = '')
	`

	var count int
	err := database.QueryRowContext(ctx, query, userID, roleName, bankID).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *entitlementRepository) Delete(ctx context.Context, entitlementID string) error {
	database := db.GetDB()
	query := `DELETE FROM entitlements WHERE entitlement_id = ?`
	_, err := database.ExecContext(ctx, query, entitlementID)
	return err
}
