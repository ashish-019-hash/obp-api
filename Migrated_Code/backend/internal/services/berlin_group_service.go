package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/ashish-019-hash/obp-api-backend/internal/models"
	"gorm.io/gorm"
)

type BerlinGroupService struct {
	db            *gorm.DB
	configService *ConfigService
}

type BerlinGroupConsent struct {
	ConsentID            string                 `json:"consentId"`
	Access               BerlinGroupAccess      `json:"access"`
	RecurringIndicator   bool                   `json:"recurringIndicator"`
	ValidUntil           time.Time              `json:"validUntil"`
	FrequencyPerDay      int                    `json:"frequencyPerDay"`
	LastActionDate       time.Time              `json:"lastActionDate"`
	ConsentStatus        string                 `json:"consentStatus"`
	UsesSoFarTodayCounter int                   `json:"usesSoFarTodayCounter"`
	Links                map[string]interface{} `json:"_links,omitempty"`
}

type BerlinGroupAccess struct {
	Accounts             []BerlinGroupAccount `json:"accounts,omitempty"`
	Balances             []BerlinGroupAccount `json:"balances,omitempty"`
	Transactions         []BerlinGroupAccount `json:"transactions,omitempty"`
	AvailableAccounts    string               `json:"availableAccounts,omitempty"`
	AvailableAccountsWithBalance string       `json:"availableAccountsWithBalance,omitempty"`
	AllPSD2              string               `json:"allPsd2,omitempty"`
}

type BerlinGroupAccount struct {
	IBAN     string `json:"iban,omitempty"`
	BBAN     string `json:"bban,omitempty"`
	PAN      string `json:"pan,omitempty"`
	MSISDN   string `json:"msisdn,omitempty"`
	Currency string `json:"currency,omitempty"`
}

func NewBerlinGroupService(db *gorm.DB, configService *ConfigService) *BerlinGroupService {
	return &BerlinGroupService{
		db:            db,
		configService: configService,
	}
}

func (bg *BerlinGroupService) CreateBerlinGroupConsent(
	userID string,
	consumerID string,
	access BerlinGroupAccess,
	recurringIndicator bool,
	validUntil time.Time,
	frequencyPerDay int,
) (*BerlinGroupConsent, error) {
	
	if !bg.configService.GetConfigBool("berlin.group.consent.enabled", false) {
		return nil, errors.New("Berlin Group consent is disabled")
	}

	consentID := fmt.Sprintf("bg_%d", time.Now().Unix())

	accessJSON, err := json.Marshal(access)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal access: %w", err)
	}

	consent := &models.Consent{
		ConsentID:         consentID,
		UserID:            userID,
		ConsumerID:        consumerID,
		Status:            "INITIATED",
		ConsentType:       "BERLIN_GROUP",
		Scopes:            string(accessJSON),
		ValidFrom:         time.Now(),
		ValidUntil:        validUntil,
		RecurringIndicator: recurringIndicator,
		FrequencyPerDay:   frequencyPerDay,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	if err := bg.db.Create(consent).Error; err != nil {
		return nil, fmt.Errorf("failed to create consent: %w", err)
	}

	bgConsent := &BerlinGroupConsent{
		ConsentID:            consentID,
		Access:               access,
		RecurringIndicator:   recurringIndicator,
		ValidUntil:           validUntil,
		FrequencyPerDay:      frequencyPerDay,
		LastActionDate:       time.Now(),
		ConsentStatus:        "received",
		UsesSoFarTodayCounter: 0,
		Links: map[string]interface{}{
			"self": map[string]string{
				"href": fmt.Sprintf("/v1/consents/%s", consentID),
			},
			"status": map[string]string{
				"href": fmt.Sprintf("/v1/consents/%s/status", consentID),
			},
		},
	}

	return bgConsent, nil
}

func (bg *BerlinGroupService) GetBerlinGroupConsent(consentID string) (*BerlinGroupConsent, error) {
	var consent models.Consent
	if err := bg.db.Where("consent_id = ? AND consent_type = ?", consentID, "BERLIN_GROUP").First(&consent).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("consent not found")
		}
		return nil, fmt.Errorf("database error: %w", err)
	}

	var access BerlinGroupAccess
	if err := json.Unmarshal([]byte(consent.Scopes), &access); err != nil {
		return nil, fmt.Errorf("failed to parse access: %w", err)
	}

	bgStatus := bg.mapStatusToBerlinGroup(consent.Status)

	bgConsent := &BerlinGroupConsent{
		ConsentID:            consent.ConsentID,
		Access:               access,
		RecurringIndicator:   consent.RecurringIndicator,
		ValidUntil:           consent.ValidUntil,
		FrequencyPerDay:      consent.FrequencyPerDay,
		LastActionDate:       consent.UpdatedAt,
		ConsentStatus:        bgStatus,
		UsesSoFarTodayCounter: consent.UsesSoFarTodayCounter,
	}

	return bgConsent, nil
}

func (bg *BerlinGroupService) UpdateBerlinGroupConsent(consentID string, usesSoFarTodayCounter int) error {
	result := bg.db.Model(&models.Consent{}).
		Where("consent_id = ? AND consent_type = ?", consentID, "BERLIN_GROUP").
		Updates(map[string]interface{}{
			"uses_so_far_today_counter": usesSoFarTodayCounter,
			"updated_at":                time.Now(),
		})

	if result.Error != nil {
		return fmt.Errorf("failed to update consent: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return errors.New("consent not found")
	}

	return nil
}

func (bg *BerlinGroupService) RevokeBerlinGroupConsent(consentID string) error {
	result := bg.db.Model(&models.Consent{}).
		Where("consent_id = ? AND consent_type = ?", consentID, "BERLIN_GROUP").
		Updates(map[string]interface{}{
			"status":     "REVOKED",
			"updated_at": time.Now(),
		})

	if result.Error != nil {
		return fmt.Errorf("failed to revoke consent: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return errors.New("consent not found")
	}

	return nil
}

func (bg *BerlinGroupService) ValidateBerlinGroupConsent(consentID string, requestedAccess string) error {
	consent, err := bg.GetBerlinGroupConsent(consentID)
	if err != nil {
		return err
	}

	if consent.ConsentStatus != "valid" {
		return fmt.Errorf("consent status is %s", consent.ConsentStatus)
	}

	if time.Now().After(consent.ValidUntil) {
		return errors.New("consent expired")
	}

	if consent.UsesSoFarTodayCounter >= consent.FrequencyPerDay {
		return errors.New("daily frequency limit exceeded")
	}


	return nil
}

func (bg *BerlinGroupService) mapStatusToBerlinGroup(status string) string {
	switch status {
	case "INITIATED":
		return "received"
	case "ACCEPTED":
		return "valid"
	case "REJECTED":
		return "rejected"
	case "REVOKED":
		return "revokedByPsu"
	case "EXPIRED":
		return "expired"
	default:
		return "received"
	}
}

func (bg *BerlinGroupService) IncrementUsageCounter(consentID string) error {
	result := bg.db.Model(&models.Consent{}).
		Where("consent_id = ? AND consent_type = ?", consentID, "BERLIN_GROUP").
		Update("uses_so_far_today_counter", gorm.Expr("uses_so_far_today_counter + 1"))

	if result.Error != nil {
		return fmt.Errorf("failed to increment usage counter: %w", result.Error)
	}

	return nil
}
