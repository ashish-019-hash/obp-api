package models

import (
	"time"
)

type Consent struct {
	ID                     int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	ConsentID              string    `json:"consent_id" gorm:"uniqueIndex;size:255;not null"`
	UserID                 string    `json:"user_id" gorm:"index;size:255;not null"`
	ConsumerID             string    `json:"consumer_id" gorm:"index;size:255"`
	BankID                 string    `json:"bank_id" gorm:"index;size:255;not null"`
	ConsentRequestID       string    `json:"consent_request_id" gorm:"size:255"`
	Status                 string    `json:"status" gorm:"size:32;not null"`
	ConsentType            string    `json:"consent_type" gorm:"size:32;not null"`
	Scopes                 string    `json:"scopes" gorm:"type:text"`
	ValidFrom              time.Time `json:"valid_from"`
	ValidUntil             time.Time `json:"valid_until"`
	RecurringIndicator     bool      `json:"recurring_indicator" gorm:"default:false"`
	FrequencyPerDay        int       `json:"frequency_per_day" gorm:"default:0"`
	UsesSoFarTodayCounter  int       `json:"uses_so_far_today_counter" gorm:"default:0"`
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"updated_at"`
	
	User               *User              `json:"user,omitempty" gorm:"foreignKey:UserID;references:UserID"`
	Bank               *Bank              `json:"bank,omitempty" gorm:"foreignKey:BankID;references:BankID"`
	TransactionRequests []TransactionRequest `json:"transaction_requests,omitempty" gorm:"foreignKey:ConsentReferenceID;references:ConsentID"`
}

func NewConsent(consentID, userID, bankID, status, consentType string) *Consent {
	return &Consent{
		ConsentID:   consentID,
		UserID:      userID,
		BankID:      bankID,
		Status:      status,
		ConsentType: consentType,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (Consent) TableName() string {
	return "consents"
}
