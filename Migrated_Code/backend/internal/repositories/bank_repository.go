package repositories

import (
	"context"
	"database/sql"

	"obp-api-backend/internal/models"
	"obp-api-backend/pkg/db"
)

type bankRepository struct {
	db *sql.DB
}

func NewBankRepository() BankRepository {
	return &bankRepository{
		db: db.GetDB(),
	}
}

func (r *bankRepository) Create(ctx context.Context, bank *models.Bank) error {
	query := `INSERT INTO banks (bank_id, short_name, full_name, logo, website, bank_routing_scheme, bank_routing_address) 
			  VALUES (?, ?, ?, ?, ?, ?, ?)`
	
	_, err := r.db.ExecContext(ctx, query,
		bank.BankId,
		bank.ShortName,
		bank.FullName,
		bank.LogoUrl,
		bank.WebsiteUrl,
		bank.BankRoutingScheme,
		bank.BankRoutingAddress,
	)
	return err
}

func (r *bankRepository) GetByID(ctx context.Context, bankID string) (*models.Bank, error) {
	query := `SELECT bank_id, short_name, full_name, logo, website, bank_routing_scheme, bank_routing_address 
			  FROM banks WHERE bank_id = ?`
	
	bank := &models.Bank{}
	
	err := r.db.QueryRowContext(ctx, query, bankID).Scan(
		&bank.BankId,
		&bank.ShortName,
		&bank.FullName,
		&bank.LogoUrl,
		&bank.WebsiteUrl,
		&bank.BankRoutingScheme,
		&bank.BankRoutingAddress,
	)
	
	if err != nil {
		return nil, err
	}
	
	return bank, nil
}

func (r *bankRepository) Update(ctx context.Context, bank *models.Bank) error {
	query := `UPDATE banks SET short_name = ?, full_name = ?, logo = ?, website = ?, 
			  bank_routing_scheme = ?, bank_routing_address = ? WHERE bank_id = ?`
	
	_, err := r.db.ExecContext(ctx, query,
		bank.ShortName,
		bank.FullName,
		bank.LogoUrl,
		bank.WebsiteUrl,
		bank.BankRoutingScheme,
		bank.BankRoutingAddress,
		bank.BankId,
	)
	return err
}

func (r *bankRepository) Delete(ctx context.Context, bankID string) error {
	query := `DELETE FROM banks WHERE bank_id = ?`
	_, err := r.db.ExecContext(ctx, query, bankID)
	return err
}

func (r *bankRepository) List(ctx context.Context, limit, offset int) ([]*models.Bank, error) {
	query := `SELECT bank_id, short_name, full_name, logo, website, bank_routing_scheme, bank_routing_address 
			  FROM banks LIMIT ? OFFSET ?`
	
	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var banks []*models.Bank
	for rows.Next() {
		bank := &models.Bank{}
		
		err := rows.Scan(
			&bank.BankId,
			&bank.ShortName,
			&bank.FullName,
			&bank.LogoUrl,
			&bank.WebsiteUrl,
			&bank.BankRoutingScheme,
			&bank.BankRoutingAddress,
		)
		if err != nil {
			return nil, err
		}
		
		banks = append(banks, bank)
	}
	
	return banks, rows.Err()
}
