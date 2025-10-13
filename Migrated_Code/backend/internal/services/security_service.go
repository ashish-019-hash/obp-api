package services

import (
	"errors"

	"github.com/ashish-019-hash/obp-api-backend/internal/models"
	"github.com/ashish-019-hash/obp-api-backend/internal/repositories"
)

type SecurityService struct {
	consentRepo     repositories.ConsentRepository
	currencyService *CurrencyService
}

func NewSecurityService(consentRepo repositories.ConsentRepository, currencyService *CurrencyService) *SecurityService {
	return &SecurityService{
		consentRepo:     consentRepo,
		currencyService: currencyService,
	}
}

func (s *SecurityService) CalculateChallengeThreshold(amount int64, currency string) (bool, error) {
	defaultThresholdUSD := int64(100000) // $1000 in cents
	
	if currency == "USD" {
		return amount >= defaultThresholdUSD, nil
	}

	decimalAmount := s.currencyService.ConvertFromSmallestUnit(amount, currency)
	usdAmount, err := s.currencyService.ConvertCurrency(decimalAmount, currency, "USD")
	if err != nil {
		return false, err
	}

	usdAmountCents := s.currencyService.ConvertToSmallestUnit(usdAmount, "USD")
	
	return usdAmountCents >= defaultThresholdUSD, nil
}

func (s *SecurityService) ValidateConsent(userID, consentType string) (*models.Consent, error) {
	consents, err := s.consentRepo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	for _, consent := range consents {
		if consent.ConsentType == consentType && consent.Status == "ACTIVE" {
			return consent, nil
		}
	}

	return nil, errors.New("no valid consent found for user and consent type")
}

func (s *SecurityService) CreateConsent(consent *models.Consent) error {
	return s.consentRepo.Create(consent)
}

func (s *SecurityService) RevokeConsent(consentID string) error {
	return s.consentRepo.UpdateStatus(consentID, "REVOKED")
}

func (s *SecurityService) GetUserConsents(userID string) ([]*models.Consent, error) {
	return s.consentRepo.GetByUserID(userID)
}

func (s *SecurityService) CheckTransactionAuthorization(userID string, amount int64, currency string) (bool, error) {
	requiresChallenge, err := s.CalculateChallengeThreshold(amount, currency)
	if err != nil {
		return false, err
	}

	if requiresChallenge {
		_, err := s.ValidateConsent(userID, "PAYMENT")
		if err != nil {
			return false, errors.New("payment consent required for high-value transaction: " + err.Error())
		}
	}

	return true, nil
}
