package models

import (
	"time"
)

type User struct {
	ID                                int64      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID                            string     `json:"user_id" gorm:"uniqueIndex;size:255;not null"`
	IDGivenByProvider                 string     `json:"id_given_by_provider" gorm:"size:255"`
	Provider                          string     `json:"provider" gorm:"size:100;not null"`
	EmailAddress                      string     `json:"email_address" gorm:"size:255;not null"`
	Name                              string     `json:"name" gorm:"size:255;not null"`
	CreatedByConsentID                *string    `json:"created_by_consent_id,omitempty" gorm:"size:255"`
	CreatedByUserInvitationID         *string    `json:"created_by_user_invitation_id,omitempty" gorm:"size:255"`
	IsOriginalUser                    bool       `json:"is_original_user"`
	IsConsentUser                     bool       `json:"is_consent_user"`
	IsDeleted                         *bool      `json:"is_deleted,omitempty"`
	LastMarketingAgreementSignedDate  *time.Time `json:"last_marketing_agreement_signed_date,omitempty"`
	LastUsedLocale                    *string    `json:"last_used_locale,omitempty" gorm:"size:10"`
	CreatedAt                         time.Time  `json:"created_at"`
	UpdatedAt                         time.Time  `json:"updated_at"`
	
	Email                             string     `json:"email" gorm:"size:255"`
	FirstName                         string     `json:"first_name" gorm:"size:255"`
	LastName                          string     `json:"last_name" gorm:"size:255"`
	ProviderID                        string     `json:"provider_id" gorm:"size:255"`
	IsActive                          bool       `json:"is_active" gorm:"default:true"`
	ConsentGiven                      bool       `json:"consent_given" gorm:"default:false"`
	
	UserCustomerLinks                 []UserCustomerLink `json:"user_customer_links,omitempty" gorm:"foreignKey:UserID;references:UserID"`
	Consents                          []Consent          `json:"consents,omitempty" gorm:"foreignKey:UserID;references:UserID"`
}

func NewUser(userID, provider, emailAddress, name string) *User {
	return &User{
		UserID:         userID,
		Provider:       provider,
		EmailAddress:   emailAddress,
		Name:           name,
		IsOriginalUser: true,
		IsConsentUser:  false,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
}

func (User) TableName() string {
	return "users"
}
