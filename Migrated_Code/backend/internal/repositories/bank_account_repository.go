package repositories

import (
	"context"
	"database/sql"
	"math/big"

	"obp-api-backend/internal/models"
	"obp-api-backend/pkg/db"
)

type BankAccountRepository interface {
	Create(ctx context.Context, account *models.BankAccount) error
	GetByID(ctx context.Context, id string) (*models.BankAccount, error)
	GetByBankID(ctx context.Context, bankID string, limit, offset int) ([]*models.BankAccount, error)
	Update(ctx context.Context, account *models.BankAccount) error
	Delete(ctx context.Context, id string) error
	GetAccountsByRouting(ctx context.Context, routingScheme, routingAddress string) ([]*models.BankAccount, error)
}

type bankAccountRepository struct {
	db *sql.DB
}

func NewBankAccountRepository() BankAccountRepository {
	return &bankAccountRepository{
		db: db.GetDB(),
	}
}

func (r *bankAccountRepository) Create(ctx context.Context, account *models.BankAccount) error {
	query := `INSERT INTO bank_accounts (account_id, bank_id, label, number, type, balance_currency, balance_amount) 
			  VALUES (?, ?, ?, ?, ?, ?, ?)`
	
	balanceStr := account.Balance.String()
	
	_, err := r.db.ExecContext(ctx, query,
		account.AccountId,
		account.BankId,
		account.Label,
		account.Number,
		account.AccountType,
		account.Currency,
		balanceStr,
	)
	return err
}

func (r *bankAccountRepository) GetByID(ctx context.Context, accountID string) (*models.BankAccount, error) {
	query := `SELECT account_id, bank_id, label, number, type, balance_currency, balance_amount 
			  FROM bank_accounts WHERE account_id = ?`
	
	account := &models.BankAccount{}
	var balanceStr string
	
	err := r.db.QueryRowContext(ctx, query, accountID).Scan(
		&account.AccountId,
		&account.BankId,
		&account.Label,
		&account.Number,
		&account.AccountType,
		&account.Currency,
		&balanceStr,
	)
	
	if err != nil {
		return nil, err
	}
	
	account.Balance, _ = new(big.Float).SetString(balanceStr)
	if account.Balance == nil {
		account.Balance = big.NewFloat(0.0)
	}
	
	return account, nil
}

func (r *bankAccountRepository)	GetByBankID(ctx context.Context, bankID string, limit, offset int) ([]*models.BankAccount, error) {
	query := `SELECT account_id, bank_id, label, number, type, balance_currency, balance_amount 
			  FROM bank_accounts WHERE bank_id = ?`
	
	rows, err := r.db.QueryContext(ctx, query, bankID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var accounts []*models.BankAccount
	for rows.Next() {
		account := &models.BankAccount{}
		var balanceStr string
		
		err := rows.Scan(
			&account.AccountId,
			&account.BankId,
			&account.Label,
			&account.Number,
			&account.AccountType,
			&account.Currency,
			&balanceStr,
		)
		if err != nil {
			return nil, err
		}
		
		account.Balance, _ = new(big.Float).SetString(balanceStr)
		if account.Balance == nil {
			account.Balance = big.NewFloat(0.0)
		}
		
		accounts = append(accounts, account)
	}
	
	return accounts, rows.Err()
}

func (r *bankAccountRepository) Update(ctx context.Context, account *models.BankAccount) error {
	query := `UPDATE bank_accounts SET label = ?, number = ?, type = ?, balance_currency = ?, balance_amount = ? WHERE account_id = ?`
	
	balanceStr := account.Balance.String()
	
	_, err := r.db.ExecContext(ctx, query,
		account.Label,
		account.Number,
		account.AccountType,
		account.Currency,
		balanceStr,
		account.AccountId,
	)
	return err
}

func (r *bankAccountRepository) Delete(ctx context.Context, accountID string) error {
	query := `DELETE FROM bank_accounts WHERE account_id = ?`
	_, err := r.db.ExecContext(ctx, query, accountID)
	return err
}

func (r *bankAccountRepository) GetAccountsByRouting(ctx context.Context, scheme, address string) ([]*models.BankAccount, error) {
	query := `SELECT account_id, bank_id, label, number, type, balance_currency, balance_amount 
			  FROM bank_accounts WHERE account_routing_scheme = ? AND account_routing_address = ?`
	
	rows, err := r.db.QueryContext(ctx, query, scheme, address)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var accounts []*models.BankAccount
	for rows.Next() {
		account := &models.BankAccount{}
		var balanceStr string
		
		err := rows.Scan(
			&account.AccountId,
			&account.BankId,
			&account.Label,
			&account.Number,
			&account.AccountType,
			&account.Currency,
			&balanceStr,
		)
		if err != nil {
			return nil, err
		}
		
		account.Balance, _ = new(big.Float).SetString(balanceStr)
		if account.Balance == nil {
			account.Balance = big.NewFloat(0.0)
		}
		
		accounts = append(accounts, account)
	}
	
	return accounts, rows.Err()
}
