package repositories

import (
	"context"
	"database/sql"
	"math/big"

	"obp-api-backend/internal/models"
	"obp-api-backend/pkg/db"
)

type counterpartyLimitRepository struct {
	db *sql.DB
}

func NewCounterpartyLimitRepository() CounterpartyLimitRepository {
	return &counterpartyLimitRepository{
		db: db.GetDB(),
	}
}

func (r *counterpartyLimitRepository) Create(ctx context.Context, limit *models.CounterpartyLimit) error {
	query := `INSERT INTO counterparty_limits (limit_id, counterparty_id, max_single_amount_currency, max_single_amount_value, max_monthly_amount_currency, max_monthly_amount_value) 
			  VALUES (?, ?, ?, ?, ?, ?)`

	maxSingleStr := limit.MaxSingleAmount.String()
	maxMonthlyStr := limit.MaxMonthlyAmount.String()

	_, err := r.db.ExecContext(ctx, query,
		limit.CounterpartyLimitId,
		limit.CounterpartyId,
		limit.Currency,
		maxSingleStr,
		limit.Currency,
		maxMonthlyStr,
	)
	return err
}

func (r *counterpartyLimitRepository) GetByCounterpartyID(ctx context.Context, counterpartyID string) (*models.CounterpartyLimit, error) {
	query := `SELECT limit_id, counterparty_id, max_single_amount_currency, max_single_amount_value, max_monthly_amount_currency, max_monthly_amount_value 
			  FROM counterparty_limits WHERE counterparty_id = ?`

	limit := &models.CounterpartyLimit{}
	var maxSingleStr, maxMonthlyStr, currency1, currency2 string

	err := r.db.QueryRowContext(ctx, query, counterpartyID).Scan(
		&limit.CounterpartyLimitId,
		&limit.CounterpartyId,
		&currency1,
		&maxSingleStr,
		&currency2,
		&maxMonthlyStr,
	)

	if err != nil {
		return nil, err
	}

	limit.Currency = currency1
	limit.MaxSingleAmount, _ = new(big.Float).SetString(maxSingleStr)
	if limit.MaxSingleAmount == nil {
		limit.MaxSingleAmount = big.NewFloat(0.0)
	}

	limit.MaxMonthlyAmount, _ = new(big.Float).SetString(maxMonthlyStr)
	if limit.MaxMonthlyAmount == nil {
		limit.MaxMonthlyAmount = big.NewFloat(0.0)
	}

	return limit, nil
}

func (r *counterpartyLimitRepository) Update(ctx context.Context, limit *models.CounterpartyLimit) error {
	query := `UPDATE counterparty_limits SET max_single_amount_currency = ?, max_single_amount_value = ?, 
			  max_monthly_amount_currency = ?, max_monthly_amount_value = ? WHERE limit_id = ?`

	maxSingleStr := limit.MaxSingleAmount.String()
	maxMonthlyStr := limit.MaxMonthlyAmount.String()

	_, err := r.db.ExecContext(ctx, query,
		limit.Currency,
		maxSingleStr,
		limit.Currency,
		maxMonthlyStr,
		limit.CounterpartyLimitId,
	)
	return err
}

func (r *counterpartyLimitRepository) Delete(ctx context.Context, limitID string) error {
	query := `DELETE FROM counterparty_limits WHERE limit_id = ?`
	_, err := r.db.ExecContext(ctx, query, limitID)
	return err
}

func (r *counterpartyLimitRepository) ValidateLimit(ctx context.Context, counterpartyID string, amount *big.Float, currency string) (bool, error) {
	limit, err := r.GetByCounterpartyID(ctx, counterpartyID)
	if err != nil {
		return false, err
	}

	if limit.Currency != currency {
		return false, nil
	}

	return amount.Cmp(limit.MaxSingleAmount) <= 0, nil
}
