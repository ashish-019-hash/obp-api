package repositories

import (
	"context"
	"database/sql"
	"math/big"
	"time"

	"obp-api-backend/internal/models"
	"obp-api-backend/pkg/db"
)

type transactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository() TransactionRepository {
	return &transactionRepository{
		db: db.GetDB(),
	}
}

func (r *transactionRepository) Create(ctx context.Context, transaction *models.Transaction) error {
	query := `INSERT INTO transactions (transaction_id, account_id, counterparty_id, amount_currency, amount_value, 
			  description, posted_date, completed_date, new_balance_currency, new_balance_amount, value_date) 
			  VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	amountStr := transaction.Amount.String()
	balanceStr := transaction.Balance.String()

	_, err := r.db.ExecContext(ctx, query,
		transaction.Id,
		transaction.ThisAccount,
		transaction.OtherAccount,
		transaction.Currency,
		amountStr,
		transaction.Description,
		transaction.StartDate,
		transaction.FinishDate,
		transaction.Currency,
		balanceStr,
		transaction.StartDate,
	)
	return err
}

func (r *transactionRepository) GetByID(ctx context.Context, transactionID string) (*models.Transaction, error) {
	query := `SELECT transaction_id, account_id, counterparty_id, amount_currency, amount_value, 
			  description, posted_date, completed_date, new_balance_currency, new_balance_amount, value_date 
			  FROM transactions WHERE transaction_id = ?`

	transaction := &models.Transaction{}
	var amountStr, balanceStr, newBalanceCurrency, valueDate string

	err := r.db.QueryRowContext(ctx, query, transactionID).Scan(
		&transaction.Id,
		&transaction.ThisAccount,
		&transaction.OtherAccount,
		&transaction.Currency,
		&amountStr,
		&transaction.Description,
		&transaction.StartDate,
		&transaction.FinishDate,
		&newBalanceCurrency,
		&balanceStr,
		&valueDate,
	)

	if err != nil {
		return nil, err
	}

	transaction.Amount, _ = new(big.Float).SetString(amountStr)
	if transaction.Amount == nil {
		transaction.Amount = big.NewFloat(0.0)
	}

	transaction.Balance, _ = new(big.Float).SetString(balanceStr)
	if transaction.Balance == nil {
		transaction.Balance = big.NewFloat(0.0)
	}

	return transaction, nil
}

func (r *transactionRepository) GetByAccountID(ctx context.Context, accountID string, limit, offset int) ([]*models.Transaction, error) {
	query := `SELECT transaction_id, account_id, counterparty_id, amount_currency, amount_value, 
			  description, posted_date, completed_date, new_balance_currency, new_balance_amount, value_date 
			  FROM transactions WHERE account_id = ? ORDER BY posted_date DESC LIMIT ? OFFSET ?`

	rows, err := r.db.QueryContext(ctx, query, accountID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []*models.Transaction
	for rows.Next() {
		transaction := &models.Transaction{}
		var amountStr, balanceStr, newBalanceCurrency, valueDate string

		err := rows.Scan(
			&transaction.Id,
			&transaction.ThisAccount,
			&transaction.OtherAccount,
			&transaction.Currency,
			&amountStr,
			&transaction.Description,
			&transaction.StartDate,
			&transaction.FinishDate,
			&newBalanceCurrency,
			&balanceStr,
			&valueDate,
		)
		if err != nil {
			return nil, err
		}

		transaction.Amount, _ = new(big.Float).SetString(amountStr)
		if transaction.Amount == nil {
			transaction.Amount = big.NewFloat(0.0)
		}

		transaction.Balance, _ = new(big.Float).SetString(balanceStr)
		if transaction.Balance == nil {
			transaction.Balance = big.NewFloat(0.0)
		}

		transactions = append(transactions, transaction)
	}

	return transactions, rows.Err()
}

func (r *transactionRepository) GetByDateRange(ctx context.Context, accountID string, fromDate, toDate time.Time) ([]*models.Transaction, error) {
	query := `SELECT transaction_id, account_id, counterparty_id, amount_currency, amount_value, 
			  description, posted_date, completed_date, new_balance_currency, new_balance_amount, value_date 
			  FROM transactions WHERE account_id = ? AND posted_date BETWEEN ? AND ? ORDER BY posted_date DESC`

	rows, err := r.db.QueryContext(ctx, query, accountID, fromDate.Format(time.RFC3339), toDate.Format(time.RFC3339))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []*models.Transaction
	for rows.Next() {
		transaction := &models.Transaction{}
		var amountStr, balanceStr, newBalanceCurrency, valueDate string

		err := rows.Scan(
			&transaction.Id,
			&transaction.ThisAccount,
			&transaction.OtherAccount,
			&transaction.Currency,
			&amountStr,
			&transaction.Description,
			&transaction.StartDate,
			&transaction.FinishDate,
			&newBalanceCurrency,
			&balanceStr,
			&valueDate,
		)
		if err != nil {
			return nil, err
		}

		transaction.Amount, _ = new(big.Float).SetString(amountStr)
		if transaction.Amount == nil {
			transaction.Amount = big.NewFloat(0.0)
		}

		transaction.Balance, _ = new(big.Float).SetString(balanceStr)
		if transaction.Balance == nil {
			transaction.Balance = big.NewFloat(0.0)
		}

		transactions = append(transactions, transaction)
	}

	return transactions, rows.Err()
}

func (r *transactionRepository) GetCreditTransactions(ctx context.Context, accountID string) ([]*models.Transaction, error) {
	query := `SELECT transaction_id, account_id, counterparty_id, amount_currency, amount_value, 
			  description, posted_date, completed_date, new_balance_currency, new_balance_amount, value_date 
			  FROM transactions WHERE account_id = ? AND CAST(amount_value AS REAL) > 0 ORDER BY posted_date DESC`

	rows, err := r.db.QueryContext(ctx, query, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []*models.Transaction
	for rows.Next() {
		transaction := &models.Transaction{}
		var amountStr, balanceStr, newBalanceCurrency, valueDate string

		err := rows.Scan(
			&transaction.Id,
			&transaction.ThisAccount,
			&transaction.OtherAccount,
			&transaction.Currency,
			&amountStr,
			&transaction.Description,
			&transaction.StartDate,
			&transaction.FinishDate,
			&newBalanceCurrency,
			&balanceStr,
			&valueDate,
		)
		if err != nil {
			return nil, err
		}

		transaction.Amount, _ = new(big.Float).SetString(amountStr)
		if transaction.Amount == nil {
			transaction.Amount = big.NewFloat(0.0)
		}

		transaction.Balance, _ = new(big.Float).SetString(balanceStr)
		if transaction.Balance == nil {
			transaction.Balance = big.NewFloat(0.0)
		}

		transactions = append(transactions, transaction)
	}

	return transactions, rows.Err()
}

func (r *transactionRepository) GetDebitTransactions(ctx context.Context, accountID string) ([]*models.Transaction, error) {
	query := `SELECT transaction_id, account_id, counterparty_id, amount_currency, amount_value, 
			  description, posted_date, completed_date, new_balance_currency, new_balance_amount, value_date 
			  FROM transactions WHERE account_id = ? AND CAST(amount_value AS REAL) < 0 ORDER BY posted_date DESC`

	rows, err := r.db.QueryContext(ctx, query, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []*models.Transaction
	for rows.Next() {
		transaction := &models.Transaction{}
		var amountStr, balanceStr, newBalanceCurrency, valueDate string

		err := rows.Scan(
			&transaction.Id,
			&transaction.ThisAccount,
			&transaction.OtherAccount,
			&transaction.Currency,
			&amountStr,
			&transaction.Description,
			&transaction.StartDate,
			&transaction.FinishDate,
			&newBalanceCurrency,
			&balanceStr,
			&valueDate,
		)
		if err != nil {
			return nil, err
		}

		transaction.Amount, _ = new(big.Float).SetString(amountStr)
		if transaction.Amount == nil {
			transaction.Amount = big.NewFloat(0.0)
		}

		transaction.Balance, _ = new(big.Float).SetString(balanceStr)
		if transaction.Balance == nil {
			transaction.Balance = big.NewFloat(0.0)
		}

		transactions = append(transactions, transaction)
	}

	return transactions, rows.Err()
}

func (r *transactionRepository) CalculateBalance(ctx context.Context, accountID string) (*big.Float, error) {
	query := `SELECT COALESCE(SUM(CAST(amount_value AS REAL)), 0) FROM transactions WHERE account_id = ?`

	var balance float64
	err := r.db.QueryRowContext(ctx, query, accountID).Scan(&balance)
	if err != nil {
		return nil, err
	}

	return big.NewFloat(balance), nil
}
