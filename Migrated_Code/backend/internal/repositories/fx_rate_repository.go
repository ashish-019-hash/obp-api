package repositories

import (
	"context"
	"database/sql"
	"math/big"

	"obp-api-backend/pkg/db"
)

type fxRateRepository struct {
	db *sql.DB
}

func NewFXRateRepository() FXRateRepository {
	return &fxRateRepository{
		db: db.GetDB(),
	}
}

func (r *fxRateRepository) Create(ctx context.Context, bankID, fromCurrency, toCurrency string, rate *big.Float) error {
	query := `INSERT INTO fx_rates (bank_id, from_currency, to_currency, rate) VALUES (?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query, bankID, fromCurrency, toCurrency, rate.String())
	return err
}

func (r *fxRateRepository) GetRate(ctx context.Context, bankID, fromCurrency, toCurrency string) (*big.Float, error) {
	query := `SELECT rate FROM fx_rates WHERE bank_id = ? AND from_currency = ? AND to_currency = ? 
			  ORDER BY created_at DESC LIMIT 1`

	var rateStr string
	err := r.db.QueryRowContext(ctx, query, bankID, fromCurrency, toCurrency).Scan(&rateStr)
	if err != nil {
		return nil, err
	}

	rate, ok := new(big.Float).SetString(rateStr)
	if !ok {
		return nil, sql.ErrNoRows
	}

	return rate, nil
}

func (r *fxRateRepository) GetLatestRate(ctx context.Context, fromCurrency, toCurrency string) (*big.Float, error) {
	query := `SELECT rate FROM fx_rates WHERE from_currency = ? AND to_currency = ? 
			  ORDER BY created_at DESC LIMIT 1`

	var rateStr string
	err := r.db.QueryRowContext(ctx, query, fromCurrency, toCurrency).Scan(&rateStr)
	if err != nil {
		return nil, err
	}

	rate, ok := new(big.Float).SetString(rateStr)
	if !ok {
		return nil, sql.ErrNoRows
	}

	return rate, nil
}
