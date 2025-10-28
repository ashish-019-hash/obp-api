package repositories

import (
	"github.com/ashish-019-hash/obp-api-backend/internal/models"
	"gorm.io/gorm"
)

type bankRepository struct {
	db *gorm.DB
}

func NewBankRepository(db *gorm.DB) BankRepository {
	return &bankRepository{db: db}
}

func (r *bankRepository) Create(bank *models.Bank) error {
	return r.db.Create(bank).Error
}

func (r *bankRepository) GetByID(id int64) (*models.Bank, error) {
	var bank models.Bank
	err := r.db.First(&bank, id).Error
	if err != nil {
		return nil, err
	}
	return &bank, nil
}

func (r *bankRepository) GetByBankID(bankID string) (*models.Bank, error) {
	var bank models.Bank
	err := r.db.Where("bank_id = ?", bankID).First(&bank).Error
	if err != nil {
		return nil, err
	}
	return &bank, nil
}

func (r *bankRepository) Update(bank *models.Bank) error {
	return r.db.Save(bank).Error
}

func (r *bankRepository) Delete(id int64) error {
	return r.db.Delete(&models.Bank{}, id).Error
}

func (r *bankRepository) List(limit, offset int) ([]*models.Bank, error) {
	var banks []*models.Bank
	err := r.db.Limit(limit).Offset(offset).Find(&banks).Error
	return banks, err
}
