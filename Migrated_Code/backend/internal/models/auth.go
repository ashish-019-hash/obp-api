package models

import (
	"time"
	"golang.org/x/crypto/bcrypt"
)

type Consumer struct {
	ID                    int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	ConsumerID            string    `json:"consumer_id" gorm:"uniqueIndex;size:255;not null"`
	ConsumerKey           string    `json:"consumer_key" gorm:"uniqueIndex;size:255;not null"`
	ConsumerSecret        string    `json:"consumer_secret" gorm:"size:255;not null"`
	Name                  string    `json:"name" gorm:"size:255;not null"`
	AppType               string    `json:"app_type" gorm:"size:50;default:'Web'"`
	Description           string    `json:"description" gorm:"type:text"`
	DeveloperEmail        string    `json:"developer_email" gorm:"size:255;not null"`
	RedirectURL           string    `json:"redirect_url" gorm:"size:500"`
	IsActive              bool      `json:"is_active" gorm:"default:true"`
	CreatedByUserID       string    `json:"created_by_user_id" gorm:"size:255"`
	ClientCertificate     string    `json:"client_certificate,omitempty" gorm:"type:text"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}

type Token struct {
	ID                    int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	TokenType             string    `json:"token_type" gorm:"size:50;not null"` // "access", "refresh", "request", "id_token"
	TokenValue            string    `json:"token_value" gorm:"uniqueIndex;size:500;not null"`
	TokenSecret           string    `json:"token_secret,omitempty" gorm:"size:255"`
	ConsumerID            string    `json:"consumer_id" gorm:"index;size:255;not null"`
	UserID                string    `json:"user_id,omitempty" gorm:"index;size:255"`
	CallbackURL           string    `json:"callback_url,omitempty" gorm:"size:500"`
	Verifier              string    `json:"verifier,omitempty" gorm:"size:255"`
	Duration              int64     `json:"duration" gorm:"default:2419200"` // 4 weeks in seconds
	ExpiresAt             time.Time `json:"expires_at"`
	IsActive              bool      `json:"is_active" gorm:"default:true"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
	
	Consumer              *Consumer `json:"consumer,omitempty" gorm:"foreignKey:ConsumerID;references:ConsumerID"`
	User                  *User     `json:"user,omitempty" gorm:"foreignKey:UserID;references:UserID"`
}

type UserCredential struct {
	ID                    int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID                string    `json:"user_id" gorm:"index;size:255;not null"`
	Username              string    `json:"username" gorm:"uniqueIndex;size:255;not null"`
	PasswordHash          string    `json:"password_hash" gorm:"size:255;not null"`
	Salt                  string    `json:"salt" gorm:"size:255;not null"`
	IsActive              bool      `json:"is_active" gorm:"default:true"`
	LastLoginAt           *time.Time `json:"last_login_at,omitempty"`
	FailedLoginAttempts   int       `json:"failed_login_attempts" gorm:"default:0"`
	LockedUntil           *time.Time `json:"locked_until,omitempty"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
	
	User                  *User     `json:"user,omitempty" gorm:"foreignKey:UserID;references:UserID"`
}

type Entitlement struct {
	ID                    int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	EntitlementID         string    `json:"entitlement_id" gorm:"uniqueIndex;size:255;not null"`
	UserID                string    `json:"user_id" gorm:"index;size:255;not null"`
	RoleName              string    `json:"role_name" gorm:"size:100;not null"`
	BankID                string    `json:"bank_id,omitempty" gorm:"index;size:255"`
	IsActive              bool      `json:"is_active" gorm:"default:true"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
	
	User                  *User     `json:"user,omitempty" gorm:"foreignKey:UserID;references:UserID"`
	Bank                  *Bank     `json:"bank,omitempty" gorm:"foreignKey:BankID;references:BankID"`
}

type LoginAttempt struct {
	ID                    int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID                string    `json:"user_id,omitempty" gorm:"index;size:255"`
	Username              string    `json:"username,omitempty" gorm:"index;size:255"`
	IPAddress             string    `json:"ip_address" gorm:"size:45;not null"`
	UserAgent             string    `json:"user_agent" gorm:"size:500"`
	AuthMethod            string    `json:"auth_method" gorm:"size:50;not null"`
	Success               bool      `json:"success" gorm:"not null"`
	FailureReason         string    `json:"failure_reason,omitempty" gorm:"size:255"`
	AttemptedAt           time.Time `json:"attempted_at"`
}

func NewConsumer(consumerKey, consumerSecret, name, developerEmail string) *Consumer {
	return &Consumer{
		ConsumerID:     generateSecureID(),
		ConsumerKey:    consumerKey,
		ConsumerSecret: consumerSecret,
		Name:           name,
		AppType:        "Web",
		DeveloperEmail: developerEmail,
		IsActive:       true,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
}

func NewToken(tokenType, tokenValue, consumerID string, duration int64) *Token {
	return &Token{
		TokenType:   tokenType,
		TokenValue:  tokenValue,
		ConsumerID:  consumerID,
		Duration:    duration,
		ExpiresAt:   time.Now().Add(time.Duration(duration) * time.Second),
		IsActive:    true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func NewUserCredential(userID, username, password string) (*UserCredential, error) {
	salt := generateSecureID()
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	
	return &UserCredential{
		UserID:       userID,
		Username:     username,
		PasswordHash: string(passwordHash),
		Salt:         salt,
		IsActive:     true,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}, nil
}

func (uc *UserCredential) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(uc.PasswordHash), []byte(password+uc.Salt))
	return err == nil
}

func (uc *UserCredential) IsLocked() bool {
	if uc.LockedUntil == nil {
		return false
	}
	return time.Now().Before(*uc.LockedUntil)
}

func NewEntitlement(userID, roleName string, bankID *string) *Entitlement {
	entitlement := &Entitlement{
		EntitlementID: generateSecureID(),
		UserID:        userID,
		RoleName:      roleName,
		IsActive:      true,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	if bankID != nil {
		entitlement.BankID = *bankID
	}
	return entitlement
}

func NewLoginAttempt(userID, username, ipAddress, userAgent, authMethod string, success bool, failureReason string) *LoginAttempt {
	return &LoginAttempt{
		UserID:        userID,
		Username:      username,
		IPAddress:     ipAddress,
		UserAgent:     userAgent,
		AuthMethod:    authMethod,
		Success:       success,
		FailureReason: failureReason,
		AttemptedAt:   time.Now(),
	}
}

func generateSecureID() string {
	return time.Now().Format("20060102150405") + generateRandomString(16)
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[time.Now().UnixNano()%int64(len(charset))]
	}
	return string(b)
}

func (Consumer) TableName() string {
	return "consumers"
}

func (Token) TableName() string {
	return "tokens"
}

func (UserCredential) TableName() string {
	return "user_credentials"
}

func (Entitlement) TableName() string {
	return "entitlements"
}

func (LoginAttempt) TableName() string {
	return "login_attempts"
}
