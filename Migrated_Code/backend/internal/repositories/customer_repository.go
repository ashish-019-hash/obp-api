package repositories

import (
	"github.com/ashish-019-hash/obp-api-backend/internal/models"
	"gorm.io/gorm"
)

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{db: db}
}

func (r *customerRepository) Create(customer *models.Customer) error {
	return r.db.Create(customer).Error
}

func (r *customerRepository) GetByID(id int64) (*models.Customer, error) {
	var customer models.Customer
	err := r.db.Preload("Bank").Preload("UserCustomerLinks").Preload("CustomerAccountLinks").First(&customer, id).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *customerRepository) GetByCustomerID(customerID string) (*models.Customer, error) {
	var customer models.Customer
	err := r.db.Preload("Bank").Preload("UserCustomerLinks").Preload("CustomerAccountLinks").Where("customer_id = ?", customerID).First(&customer).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *customerRepository) GetByBankID(bankID string) ([]*models.Customer, error) {
	var customers []*models.Customer
	err := r.db.Preload("Bank").Where("bank_id = ?", bankID).Find(&customers).Error
	return customers, err
}

func (r *customerRepository) UpdateKYCStatus(customerID string, status bool) error {
	return r.db.Model(&models.Customer{}).Where("customer_id = ?", customerID).Update("kyc_status", status).Error
}

func (r *customerRepository) Update(customer *models.Customer) error {
	return r.db.Save(customer).Error
}

func (r *customerRepository) Delete(id int64) error {
	return r.db.Delete(&models.Customer{}, id).Error
}

func (r *customerRepository) List(limit, offset int) ([]*models.Customer, error) {
	var customers []*models.Customer
	err := r.db.Preload("Bank").Limit(limit).Offset(offset).Find(&customers).Error
	return customers, err
}
