package models

import (
	"time"
)

type Scope struct {
	ID         int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	ScopeID    string    `json:"scope_id" gorm:"uniqueIndex;size:255;not null"`
	BankID     string    `json:"bank_id" gorm:"index;size:255"`
	ConsumerID string    `json:"consumer_id" gorm:"index;size:255;not null"`
	RoleName   string    `json:"role_name" gorm:"size:100;not null"`
	IsActive   bool      `json:"is_active" gorm:"default:true"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (Scope) TableName() string {
	return "scopes"
}

func NewScope(consumerID, roleName string, bankID *string) *Scope {
	scope := &Scope{
		ScopeID:    generateSecureID(),
		ConsumerID: consumerID,
		RoleName:   roleName,
		IsActive:   true,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	if bankID != nil {
		scope.BankID = *bankID
	}
	return scope
}

type ViewPermission struct {
	ID             int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	PermissionID   string    `json:"permission_id" gorm:"uniqueIndex;size:255;not null"`
	ViewID         string    `json:"view_id" gorm:"index;size:255;not null"`
	BankID         string    `json:"bank_id" gorm:"index;size:255"`
	AccountID      string    `json:"account_id" gorm:"index;size:255"`
	PermissionName string    `json:"permission_name" gorm:"size:255;not null"`
	ExtraData      string    `json:"extra_data" gorm:"type:text"`
	IsActive       bool      `json:"is_active" gorm:"default:true"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (ViewPermission) TableName() string {
	return "view_permissions"
}

func NewViewPermission(viewID, permissionName string, bankID, accountID *string) *ViewPermission {
	permission := &ViewPermission{
		PermissionID:   generateSecureID(),
		ViewID:         viewID,
		PermissionName: permissionName,
		IsActive:       true,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	if bankID != nil {
		permission.BankID = *bankID
	}
	if accountID != nil {
		permission.AccountID = *accountID
	}
	return permission
}

type UserAuthContext struct {
	ID         int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	ContextID  string    `json:"context_id" gorm:"uniqueIndex;size:255;not null"`
	UserID     string    `json:"user_id" gorm:"index;size:255;not null"`
	ConsumerID string    `json:"consumer_id" gorm:"index;size:255"`
	Key        string    `json:"key" gorm:"size:255;not null"`
	Value      string    `json:"value" gorm:"type:text;not null"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (UserAuthContext) TableName() string {
	return "user_auth_contexts"
}

func NewUserAuthContext(userID, consumerID, key, value string) *UserAuthContext {
	return &UserAuthContext{
		ContextID:  generateSecureID(),
		UserID:     userID,
		ConsumerID: consumerID,
		Key:        key,
		Value:      value,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}

type ConsentAuthContext struct {
	ID        int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	ContextID string    `json:"context_id" gorm:"uniqueIndex;size:255;not null"`
	ConsentID string    `json:"consent_id" gorm:"index;size:255;not null"`
	Key       string    `json:"key" gorm:"size:255;not null"`
	Value     string    `json:"value" gorm:"type:text;not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (ConsentAuthContext) TableName() string {
	return "consent_auth_contexts"
}

func NewConsentAuthContext(consentID, key, value string) *ConsentAuthContext {
	return &ConsentAuthContext{
		ContextID: generateSecureID(),
		ConsentID: consentID,
		Key:       key,
		Value:     value,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

type UserLock struct {
	ID           int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	LockID       string    `json:"lock_id" gorm:"uniqueIndex;size:255;not null"`
	UserID       string    `json:"user_id" gorm:"index;size:255;not null"`
	TypeOfLock   string    `json:"type_of_lock" gorm:"size:100;not null"`
	LockReason   string    `json:"lock_reason" gorm:"type:text"`
	LastLockDate time.Time `json:"last_lock_date"`
	IsActive     bool      `json:"is_active" gorm:"default:true"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (UserLock) TableName() string {
	return "user_locks"
}

func NewUserLock(userID, typeOfLock, reason string) *UserLock {
	return &UserLock{
		LockID:       generateSecureID(),
		UserID:       userID,
		TypeOfLock:   typeOfLock,
		LockReason:   reason,
		LastLockDate: time.Now(),
		IsActive:     true,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

type AuthenticationTypeValidation struct {
	ID          int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	ValidationID string   `json:"validation_id" gorm:"uniqueIndex;size:255;not null"`
	OperationID string    `json:"operation_id" gorm:"index;size:255;not null"`
	AuthTypes   string    `json:"auth_types" gorm:"type:text;not null"`
	IsActive    bool      `json:"is_active" gorm:"default:true"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (AuthenticationTypeValidation) TableName() string {
	return "authentication_type_validations"
}

func NewAuthenticationTypeValidation(operationID string, authTypes []string) *AuthenticationTypeValidation {
	authTypesStr := ""
	for i, authType := range authTypes {
		if i > 0 {
			authTypesStr += ","
		}
		authTypesStr += authType
	}
	
	return &AuthenticationTypeValidation{
		ValidationID: generateSecureID(),
		OperationID:  operationID,
		AuthTypes:    authTypesStr,
		IsActive:     true,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

func (atv *AuthenticationTypeValidation) GetAuthTypes() []string {
	if atv.AuthTypes == "" {
		return []string{}
	}
	authTypes := []string{}
	for _, authType := range []string{"JWT", "OAuth1", "OAuth2", "DirectLogin", "GatewayLogin"} {
		if contains(atv.AuthTypes, authType) {
			authTypes = append(authTypes, authType)
		}
	}
	return authTypes
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || 
		(len(s) > len(substr) && (s[:len(substr)+1] == substr+"," || 
		s[len(s)-len(substr)-1:] == ","+substr || 
		len(s) > len(substr)*2 && s[len(s)-len(substr)-1:len(s)-len(substr)] == ","+substr+",")))
}
