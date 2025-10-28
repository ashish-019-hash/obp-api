package repositories

import (
	"github.com/ashish-019-hash/obp-api-backend/internal/models"
	"gorm.io/gorm"
)

type bankAccountRepository struct {
	db *gorm.DB
}

func NewBankAccountRepository(db *gorm.DB) BankAccountRepository {
	return &bankAccountRepository{db: db}
}

func (r *bankAccountRepository) Create(account *models.BankAccount) error {
	return r.db.Create(account).Error
}

func (r *bankAccountRepository) GetByID(id int64) (*models.BankAccount, error) {
	var account models.BankAccount
	err := r.db.Preload("Bank").Preload("AccountRoutings").First(&account, id).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *bankAccountRepository) GetByAccountID(accountID string) (*models.BankAccount, error) {
	var account models.BankAccount
	err := r.db.Preload("Bank").Preload("AccountRoutings").Where("account_id = ?", accountID).First(&account).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *bankAccountRepository) GetByBankID(bankID string) ([]*models.BankAccount, error) {
	var accounts []*models.BankAccount
	err := r.db.Preload("Bank").Where("bank_id = ?", bankID).Find(&accounts).Error
	return accounts, err
}

func (r *bankAccountRepository) UpdateBalance(accountID string, newBalance int64) error {
	return r.db.Model(&models.BankAccount{}).Where("account_id = ?", accountID).Update("balance", newBalance).Error
}

func (r *bankAccountRepository) Update(account *models.BankAccount) error {
	return r.db.Save(account).Error
}

func (r *bankAccountRepository) Delete(id int64) error {
	return r.db.Delete(&models.BankAccount{}, id).Error
}

func (r *bankAccountRepository) List(limit, offset int) ([]*models.BankAccount, error) {
	var accounts []*models.BankAccount
	err := r.db.Preload("Bank").Limit(limit).Offset(offset).Find(&accounts).Error
	return accounts, err
}
