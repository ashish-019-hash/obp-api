package services

import (
	"context"
	"obp-api-backend/internal/models"
	"obp-api-backend/internal/repositories"
)

type CustomerService interface {
	GetCustomersByBankID(ctx context.Context, bankID string, limit, offset int) ([]*models.Customer, error)
	GetCustomerByID(ctx context.Context, customerID string) (*models.Customer, error)
	CreateCustomer(ctx context.Context, customer *models.Customer) error
}

type customerService struct {
	customerRepo repositories.CustomerRepository
}

func NewCustomerService(customerRepo repositories.CustomerRepository) CustomerService {
	return &customerService{
		customerRepo: customerRepo,
	}
}

func (s *customerService) GetCustomersByBankID(ctx context.Context, bankID string, limit, offset int) ([]*models.Customer, error) {
	return s.customerRepo.GetByBankID(ctx, bankID)
}

func (s *customerService) GetCustomerByID(ctx context.Context, customerID string) (*models.Customer, error) {
	return s.customerRepo.GetByID(ctx, customerID)
}

func (s *customerService) CreateCustomer(ctx context.Context, customer *models.Customer) error {
	return s.customerRepo.Create(ctx, customer)
}
