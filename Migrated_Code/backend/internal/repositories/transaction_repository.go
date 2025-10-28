package repositories

import (
	"github.com/ashish-019-hash/obp-api-backend/internal/models"
	"gorm.io/gorm"
)

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) Create(transaction *models.Transaction) error {
	return r.db.Create(transaction).Error
}

func (r *transactionRepository) GetByID(id int64) (*models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("BankAccount").Preload("Counterparty").First(&transaction, id).Error
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (r *transactionRepository) GetByTransactionID(transactionID string) (*models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("BankAccount").Preload("Counterparty").Where("transaction_id = ?", transactionID).First(&transaction).Error
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (r *transactionRepository) GetByAccountID(accountID string) ([]*models.Transaction, error) {
	var transactions []*models.Transaction
	err := r.db.Preload("BankAccount").Preload("Counterparty").Where("account_id = ?", accountID).Order("created_at DESC").Find(&transactions).Error
	return transactions, err
}

func (r *transactionRepository) GetByDateRange(accountID string, startDate, endDate string) ([]*models.Transaction, error) {
	var transactions []*models.Transaction
	err := r.db.Preload("BankAccount").Preload("Counterparty").
		Where("account_id = ? AND created_at BETWEEN ? AND ?", accountID, startDate, endDate).
		Order("created_at DESC").Find(&transactions).Error
	return transactions, err
}

func (r *transactionRepository) Update(transaction *models.Transaction) error {
	return r.db.Save(transaction).Error
}

func (r *transactionRepository) Delete(id int64) error {
	return r.db.Delete(&models.Transaction{}, id).Error
}

func (r *transactionRepository) List(limit, offset int) ([]*models.Transaction, error) {
	var transactions []*models.Transaction
	err := r.db.Preload("BankAccount").Preload("Counterparty").Limit(limit).Offset(offset).Order("created_at DESC").Find(&transactions).Error
	return transactions, err
}
