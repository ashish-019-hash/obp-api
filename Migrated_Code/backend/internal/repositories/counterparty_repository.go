package repositories

import (
	"context"
	"database/sql"

	"obp-api-backend/internal/models"
	"obp-api-backend/pkg/db"
)

type counterpartyRepository struct {
	db *sql.DB
}

func NewCounterpartyRepository() CounterpartyRepository {
	return &counterpartyRepository{
		db: db.GetDB(),
	}
}

func (r *counterpartyRepository) Create(ctx context.Context, counterparty *models.Counterparty) error {
	query := `INSERT INTO counterparties (counterparty_id, name, created_by_user_id, this_bank_id, this_account_id, this_view_id, counterparty_id_value, kind) 
			  VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query,
		counterparty.CounterpartyId,
		counterparty.CounterpartyName,
		"", // created_by_user_id - not available in model
		counterparty.ThisBankId,
		counterparty.ThisAccountId,
		"", // this_view_id - not available in model
		"", // counterparty_id_value - not available in model
		"", // kind - not available in model
	)
	return err
}

func (r *counterpartyRepository) GetByID(ctx context.Context, counterpartyID string) (*models.Counterparty, error) {
	query := `SELECT counterparty_id, name, created_by_user_id, this_bank_id, this_account_id, this_view_id, counterparty_id_value, kind 
			  FROM counterparties WHERE counterparty_id = ?`

	counterparty := &models.Counterparty{}
	var createdByUserId, thisViewId, counterpartyIdValue, kind string

	err := r.db.QueryRowContext(ctx, query, counterpartyID).Scan(
		&counterparty.CounterpartyId,
		&counterparty.CounterpartyName,
		&createdByUserId,
		&counterparty.ThisBankId,
		&counterparty.ThisAccountId,
		&thisViewId,
		&counterpartyIdValue,
		&kind,
	)

	return counterparty, err
}

func (r *counterpartyRepository) Update(ctx context.Context, counterparty *models.Counterparty) error {
	query := `UPDATE counterparties SET name = ?, this_bank_id = ?, this_account_id = ? WHERE counterparty_id = ?`

	_, err := r.db.ExecContext(ctx, query,
		counterparty.CounterpartyName,
		counterparty.ThisBankId,
		counterparty.ThisAccountId,
		counterparty.CounterpartyId,
	)
	return err
}

func (r *counterpartyRepository) Delete(ctx context.Context, counterpartyID string) error {
	query := `DELETE FROM counterparties WHERE counterparty_id = ?`
	_, err := r.db.ExecContext(ctx, query, counterpartyID)
	return err
}

func (r *counterpartyRepository) GetByAccountID(ctx context.Context, accountID string) ([]*models.Counterparty, error) {
	query := `SELECT counterparty_id, name, created_by_user_id, this_bank_id, this_account_id, this_view_id, counterparty_id_value, kind 
			  FROM counterparties WHERE this_account_id = ?`

	rows, err := r.db.QueryContext(ctx, query, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var counterparties []*models.Counterparty
	for rows.Next() {
		counterparty := &models.Counterparty{}
		var createdByUserId, thisViewId, counterpartyIdValue, kind string

		err := rows.Scan(
			&counterparty.CounterpartyId,
			&counterparty.CounterpartyName,
			&createdByUserId,
			&counterparty.ThisBankId,
			&counterparty.ThisAccountId,
			&thisViewId,
			&counterpartyIdValue,
			&kind,
		)
		if err != nil {
			return nil, err
		}

		counterparties = append(counterparties, counterparty)
	}

	return counterparties, rows.Err()
}
