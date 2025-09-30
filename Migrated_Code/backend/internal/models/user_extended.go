package models

import (
	"time"
)

type UserAgreement struct {
	ID                int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID            string    `json:"user_id" gorm:"index;size:255;not null"`
	AgreementType     string    `json:"agreement_type" gorm:"size:100;not null"` // "terms_and_conditions", "privacy_policy", "marketing_consent"
	AgreementVersion  string    `json:"agreement_version" gorm:"size:50;not null"`
	AcceptedAt        time.Time `json:"accepted_at"`
	IsActive          bool      `json:"is_active" gorm:"default:true"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	
	User              *User     `json:"user,omitempty" gorm:"foreignKey:UserID;references:UserID"`
}

type UserInvitation struct {
	ID                int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	InvitationID      string    `json:"invitation_id" gorm:"uniqueIndex;size:255;not null"`
	Email             string    `json:"email" gorm:"index;size:255;not null"`
	FirstName         string    `json:"first_name" gorm:"size:255"`
	LastName          string    `json:"last_name" gorm:"size:255"`
	Company           string    `json:"company" gorm:"size:255"`
	InvitedByUserID   string    `json:"invited_by_user_id" gorm:"index;size:255;not null"`
	Status            string    `json:"status" gorm:"size:50;default:'PENDING'"` // "PENDING", "ACCEPTED", "EXPIRED"
	InvitationToken   string    `json:"invitation_token" gorm:"size:500;not null"`
	ExpiresAt         time.Time `json:"expires_at"`
	AcceptedAt        *time.Time `json:"accepted_at,omitempty"`
	CreatedUserID     string    `json:"created_user_id,omitempty" gorm:"size:255"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	
	InvitedByUser     *User     `json:"invited_by_user,omitempty" gorm:"foreignKey:InvitedByUserID;references:UserID"`
	CreatedUser       *User     `json:"created_user,omitempty" gorm:"foreignKey:CreatedUserID;references:UserID"`
}

type UserAttribute struct {
	ID                int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	UserAttributeID   string    `json:"user_attribute_id" gorm:"uniqueIndex;size:255;not null"`
	UserID            string    `json:"user_id" gorm:"index;size:255;not null"`
	Name              string    `json:"name" gorm:"size:255;not null"`
	Type              string    `json:"type" gorm:"size:50;not null"` // "STRING", "INTEGER", "DOUBLE", "DATE", "BOOLEAN"
	Value             string    `json:"value" gorm:"type:text;not null"`
	IsPersonal        bool      `json:"is_personal" gorm:"default:false"`
	IsActive          bool      `json:"is_active" gorm:"default:true"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	
	User              *User     `json:"user,omitempty" gorm:"foreignKey:UserID;references:UserID"`
}

type UserRefresh struct {
	ID                int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	RefreshID         string    `json:"refresh_id" gorm:"uniqueIndex;size:255;not null"`
	UserID            string    `json:"user_id" gorm:"index;size:255;not null"`
	RefreshToken      string    `json:"refresh_token" gorm:"uniqueIndex;size:500;not null"`
	ExpiresAt         time.Time `json:"expires_at"`
	IsActive          bool      `json:"is_active" gorm:"default:true"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	
	User              *User     `json:"user,omitempty" gorm:"foreignKey:UserID;references:UserID"`
}

type APISession struct {
	ID                int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	SessionID         string    `json:"session_id" gorm:"uniqueIndex;size:255;not null"`
	UserID            string    `json:"user_id" gorm:"index;size:255;not null"`
	ConsumerID        string    `json:"consumer_id" gorm:"index;size:255;not null"`
	IPAddress         string    `json:"ip_address" gorm:"size:45"`
	UserAgent         string    `json:"user_agent" gorm:"size:500"`
	LastAccessAt      time.Time `json:"last_access_at"`
	ExpiresAt         time.Time `json:"expires_at"`
	IsActive          bool      `json:"is_active" gorm:"default:true"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	
	User              *User     `json:"user,omitempty" gorm:"foreignKey:UserID;references:UserID"`
	Consumer          *Consumer `json:"consumer,omitempty" gorm:"foreignKey:ConsumerID;references:ConsumerID"`
}

type AccountWebhook struct {
	ID                int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	WebhookID         string    `json:"webhook_id" gorm:"uniqueIndex;size:255;not null"`
	BankID            string    `json:"bank_id" gorm:"index;size:255;not null"`
	AccountID         string    `json:"account_id" gorm:"index;size:255;not null"`
	UserID            string    `json:"user_id" gorm:"index;size:255;not null"`
	TriggerName       string    `json:"trigger_name" gorm:"size:100;not null"` // "OnCreateTransaction", "OnUpdateAccount", etc.
	URL               string    `json:"url" gorm:"size:500;not null"`
	HTTPMethod        string    `json:"http_method" gorm:"size:10;default:'POST'"`
	HTTPProtocol      string    `json:"http_protocol" gorm:"size:10;default:'HTTP/1.1'"`
	IsActive          bool      `json:"is_active" gorm:"default:true"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	
	User              *User     `json:"user,omitempty" gorm:"foreignKey:UserID;references:UserID"`
	Bank              *Bank     `json:"bank,omitempty" gorm:"foreignKey:BankID;references:BankID"`
	Account           *BankAccount  `json:"account,omitempty" gorm:"foreignKey:AccountID;references:AccountID"`
}

func NewUserAgreement(userID, agreementType, version string) *UserAgreement {
	return &UserAgreement{
		UserID:           userID,
		AgreementType:    agreementType,
		AgreementVersion: version,
		AcceptedAt:       time.Now(),
		IsActive:         true,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}
}

func NewUserInvitation(email, firstName, lastName, company, invitedByUserID string) *UserInvitation {
	return &UserInvitation{
		InvitationID:     generateSecureID(),
		Email:            email,
		FirstName:        firstName,
		LastName:         lastName,
		Company:          company,
		InvitedByUserID:  invitedByUserID,
		Status:           "PENDING",
		InvitationToken:  generateSecureToken(32),
		ExpiresAt:        time.Now().Add(7 * 24 * time.Hour), // 7 days
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}
}

func NewUserAttribute(userID, name, attributeType, value string, isPersonal bool) *UserAttribute {
	return &UserAttribute{
		UserAttributeID: generateSecureID(),
		UserID:          userID,
		Name:            name,
		Type:            attributeType,
		Value:           value,
		IsPersonal:      isPersonal,
		IsActive:        true,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
}

func NewUserRefresh(userID string, expirationHours int) *UserRefresh {
	return &UserRefresh{
		RefreshID:    generateSecureID(),
		UserID:       userID,
		RefreshToken: generateSecureToken(64),
		ExpiresAt:    time.Now().Add(time.Duration(expirationHours) * time.Hour),
		IsActive:     true,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

func NewAPISession(userID, consumerID, ipAddress, userAgent string, timeoutMinutes int) *APISession {
	return &APISession{
		SessionID:    generateSecureID(),
		UserID:       userID,
		ConsumerID:   consumerID,
		IPAddress:    ipAddress,
		UserAgent:    userAgent,
		LastAccessAt: time.Now(),
		ExpiresAt:    time.Now().Add(time.Duration(timeoutMinutes) * time.Minute),
		IsActive:     true,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

func NewAccountWebhook(bankID, accountID, userID, triggerName, url string) *AccountWebhook {
	return &AccountWebhook{
		WebhookID:    generateSecureID(),
		BankID:       bankID,
		AccountID:    accountID,
		UserID:       userID,
		TriggerName:  triggerName,
		URL:          url,
		HTTPMethod:   "POST",
		HTTPProtocol: "HTTP/1.1",
		IsActive:     true,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

func generateSecureToken(length int) string {
	return generateRandomString(length)
}

func (UserAgreement) TableName() string {
	return "user_agreements"
}

func (UserInvitation) TableName() string {
	return "user_invitations"
}

func (UserAttribute) TableName() string {
	return "user_attributes"
}

func (UserRefresh) TableName() string {
	return "user_refreshes"
}

func (APISession) TableName() string {
	return "api_sessions"
}

func (AccountWebhook) TableName() string {
	return "account_webhooks"
}
