package services

import (
	"context"
	"obp-api-backend/internal/models"
	"obp-api-backend/internal/repositories"
)

type BankService interface {
	GetBanks(ctx context.Context, limit, offset int) ([]*models.Bank, error)
	GetBankByID(ctx context.Context, bankID string) (*models.Bank, error)
	CreateBank(ctx context.Context, bank *models.Bank) error
}

type bankService struct {
	bankRepo repositories.BankRepository
}

func NewBankService(bankRepo repositories.BankRepository) BankService {
	return &bankService{
		bankRepo: bankRepo,
	}
}

func (s *bankService) GetBanks(ctx context.Context, limit, offset int) ([]*models.Bank, error) {
	bank, err := s.bankRepo.GetByID(ctx, "")
	if err != nil {
		return nil, err
	}
	return []*models.Bank{bank}, nil
}

func (s *bankService) GetBankByID(ctx context.Context, bankID string) (*models.Bank, error) {
	return s.bankRepo.GetByID(ctx, bankID)
}

func (s *bankService) CreateBank(ctx context.Context, bank *models.Bank) error {
	return s.bankRepo.Create(ctx, bank)
}
