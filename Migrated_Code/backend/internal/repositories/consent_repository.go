package repositories

import (
	"github.com/ashish-019-hash/obp-api-backend/internal/models"
	"gorm.io/gorm"
)

type consentRepository struct {
	db *gorm.DB
}

func NewConsentRepository(db *gorm.DB) ConsentRepository {
	return &consentRepository{db: db}
}

func (r *consentRepository) Create(consent *models.Consent) error {
	return r.db.Create(consent).Error
}

func (r *consentRepository) GetByID(id int64) (*models.Consent, error) {
	var consent models.Consent
	err := r.db.First(&consent, id).Error
	if err != nil {
		return nil, err
	}
	return &consent, nil
}

func (r *consentRepository) GetByConsentID(consentID string) (*models.Consent, error) {
	var consent models.Consent
	err := r.db.Where("consent_id = ?", consentID).First(&consent).Error
	if err != nil {
		return nil, err
	}
	return &consent, nil
}

func (r *consentRepository) GetByUserID(userID string) ([]*models.Consent, error) {
	var consents []*models.Consent
	err := r.db.Where("user_id = ?", userID).Order("created_at DESC").Find(&consents).Error
	return consents, err
}

func (r *consentRepository) UpdateStatus(consentID string, status string) error {
	return r.db.Model(&models.Consent{}).Where("consent_id = ?", consentID).Update("status", status).Error
}

func (r *consentRepository) Update(consent *models.Consent) error {
	return r.db.Save(consent).Error
}

func (r *consentRepository) Delete(id int64) error {
	return r.db.Delete(&models.Consent{}, id).Error
}

func (r *consentRepository) List(limit, offset int) ([]*models.Consent, error) {
	var consents []*models.Consent
	err := r.db.Limit(limit).Offset(offset).Order("created_at DESC").Find(&consents).Error
	return consents, err
}
