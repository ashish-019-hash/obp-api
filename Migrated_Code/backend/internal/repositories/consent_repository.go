package repositories

import (
	"context"
	"database/sql"
	"obp-api-backend/internal/models"
	"obp-api-backend/pkg/db"
)

type ConsentRepository interface {
	Create(ctx context.Context, consent *models.Consent) error
	GetByID(ctx context.Context, id string) (*models.Consent, error)
	GetByBankID(ctx context.Context, bankID string) ([]*models.Consent, error)
	Update(ctx context.Context, id string, consent *models.Consent) error
	Delete(ctx context.Context, id string) error
}

type consentRepository struct {
	db *sql.DB
}

func NewConsentRepository() ConsentRepository {
	return &consentRepository{db: db.GetDB()}
}

func (r *consentRepository) Create(ctx context.Context, consent *models.Consent) error {
	return nil
}

func (r *consentRepository) GetByID(ctx context.Context, id string) (*models.Consent, error) {
	return &models.Consent{ConsentId: id}, nil
}

func (r *consentRepository) GetByBankID(ctx context.Context, bankID string) ([]*models.Consent, error) {
	return []*models.Consent{}, nil
}

func (r *consentRepository) Update(ctx context.Context, id string, consent *models.Consent) error {
	return nil
}

func (r *consentRepository) Delete(ctx context.Context, id string) error {
	return nil
}
