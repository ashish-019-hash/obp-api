package services

import (
	"context"
	"obp-api-backend/internal/models"
	"obp-api-backend/internal/repositories"
)

type ConsentService interface {
	CreateConsent(ctx context.Context, consent *models.Consent) error
	GetConsentsByBankID(ctx context.Context, bankID string) ([]*models.Consent, error)
	GetConsentByID(ctx context.Context, consentID string) (*models.Consent, error)
}

type consentService struct {
	consentRepo repositories.ConsentRepository
}

func NewConsentService(consentRepo repositories.ConsentRepository) ConsentService {
	return &consentService{
		consentRepo: consentRepo,
	}
}

func (s *consentService) CreateConsent(ctx context.Context, consent *models.Consent) error {
	return s.consentRepo.Create(ctx, consent)
}

func (s *consentService) GetConsentsByBankID(ctx context.Context, bankID string) ([]*models.Consent, error) {
	return s.consentRepo.GetByBankID(ctx, bankID)
}

func (s *consentService) GetConsentByID(ctx context.Context, consentID string) (*models.Consent, error) {
	return s.consentRepo.GetByID(ctx, consentID)
}
