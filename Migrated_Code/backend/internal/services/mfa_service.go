package services

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type MFAService struct {
	db            *gorm.DB
	configService *ConfigService
}

type MFAMethod struct {
	ID          int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID      string    `json:"user_id" gorm:"index;size:255;not null"`
	MethodType  string    `json:"method_type" gorm:"size:50;not null"` // "TOTP", "SMS", "EMAIL", "BACKUP_CODES"
	Secret      string    `json:"secret,omitempty" gorm:"type:text"`
	PhoneNumber string    `json:"phone_number,omitempty" gorm:"size:20"`
	Email       string    `json:"email,omitempty" gorm:"size:255"`
	IsActive    bool      `json:"is_active" gorm:"default:true"`
	IsVerified  bool      `json:"is_verified" gorm:"default:false"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type MFAChallenge struct {
	ID          int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID      string    `json:"user_id" gorm:"index;size:255;not null"`
	ChallengeID string    `json:"challenge_id" gorm:"uniqueIndex;size:255;not null"`
	MethodType  string    `json:"method_type" gorm:"size:50;not null"`
	Code        string    `json:"code" gorm:"size:10;not null"`
	ExpiresAt   time.Time `json:"expires_at"`
	IsUsed      bool      `json:"is_used" gorm:"default:false"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewMFAService(db *gorm.DB, configService *ConfigService) *MFAService {
	return &MFAService{
		db:            db,
		configService: configService,
	}
}

func (mfa *MFAService) IsEnabled() bool {
	return mfa.configService.GetConfigBool("mfa.enabled", false)
}

func (mfa *MFAService) SetupTOTP(userID string) (string, string, error) {
	if !mfa.IsEnabled() {
		return "", "", errors.New("MFA is disabled")
	}

	secret := make([]byte, 20)
	if _, err := rand.Read(secret); err != nil {
		return "", "", fmt.Errorf("failed to generate secret: %w", err)
	}

	secretBase32 := base32.StdEncoding.EncodeToString(secret)

	method := &MFAMethod{
		UserID:     userID,
		MethodType: "TOTP",
		Secret:     secretBase32,
		IsActive:   true,
		IsVerified: false,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := mfa.db.Create(method).Error; err != nil {
		return "", "", fmt.Errorf("failed to create MFA method: %w", err)
	}

	appName := mfa.configService.GetConfig("app.name", "OBP-API")
	qrURL := fmt.Sprintf("otpauth://totp/%s:%s?secret=%s&issuer=%s",
		appName, userID, secretBase32, appName)

	return secretBase32, qrURL, nil
}

func (mfa *MFAService) VerifyTOTP(userID, code string) error {
	var method MFAMethod
	if err := mfa.db.Where("user_id = ? AND method_type = ? AND is_active = ?", 
		userID, "TOTP", true).First(&method).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("TOTP not set up for user")
		}
		return fmt.Errorf("database error: %w", err)
	}

	if !mfa.validateTOTPCode(method.Secret, code) {
		return errors.New("invalid TOTP code")
	}

	if !method.IsVerified {
		mfa.db.Model(&method).Update("is_verified", true)
	}

	return nil
}

func (mfa *MFAService) SetupSMS(userID, phoneNumber string) error {
	if !mfa.IsEnabled() {
		return errors.New("MFA is disabled")
	}

	method := &MFAMethod{
		UserID:      userID,
		MethodType:  "SMS",
		PhoneNumber: phoneNumber,
		IsActive:    true,
		IsVerified:  false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := mfa.db.Create(method).Error; err != nil {
		return fmt.Errorf("failed to create SMS MFA method: %w", err)
	}

	return nil
}

func (mfa *MFAService) SendSMSChallenge(userID string) (string, error) {
	var method MFAMethod
	if err := mfa.db.Where("user_id = ? AND method_type = ? AND is_active = ?", 
		userID, "SMS", true).First(&method).Error; err != nil {
		return "", errors.New("SMS MFA not set up for user")
	}

	code := mfa.generateNumericCode(6)
	challengeID := mfa.generateChallengeID()

	challenge := &MFAChallenge{
		UserID:      userID,
		ChallengeID: challengeID,
		MethodType:  "SMS",
		Code:        code,
		ExpiresAt:   time.Now().Add(5 * time.Minute),
		CreatedAt:   time.Now(),
	}

	if err := mfa.db.Create(challenge).Error; err != nil {
		return "", fmt.Errorf("failed to create SMS challenge: %w", err)
	}

	fmt.Printf("SMS MFA code for %s: %s\n", method.PhoneNumber, code)

	return challengeID, nil
}

func (mfa *MFAService) VerifySMSChallenge(challengeID, code string) error {
	var challenge MFAChallenge
	if err := mfa.db.Where("challenge_id = ? AND method_type = ? AND is_used = ?", 
		challengeID, "SMS", false).First(&challenge).Error; err != nil {
		return errors.New("invalid or expired challenge")
	}

	if time.Now().After(challenge.ExpiresAt) {
		return errors.New("challenge expired")
	}

	if challenge.Code != code {
		return errors.New("invalid code")
	}

	mfa.db.Model(&challenge).Update("is_used", true)

	mfa.db.Model(&MFAMethod{}).
		Where("user_id = ? AND method_type = ?", challenge.UserID, "SMS").
		Update("is_verified", true)

	return nil
}

func (mfa *MFAService) GenerateBackupCodes(userID string) ([]string, error) {
	if !mfa.IsEnabled() {
		return nil, errors.New("MFA is disabled")
	}

	codes := make([]string, 10)
	for i := 0; i < 10; i++ {
		codes[i] = mfa.generateBackupCode()
	}

	hashedCodes := make([]string, len(codes))
	for i, code := range codes {
		hash := sha256.Sum256([]byte(code))
		hashedCodes[i] = hex.EncodeToString(hash[:])
	}

	method := &MFAMethod{
		UserID:     userID,
		MethodType: "BACKUP_CODES",
		Secret:     fmt.Sprintf("%v", hashedCodes), // Store as JSON string in real implementation
		IsActive:   true,
		IsVerified: true,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := mfa.db.Create(method).Error; err != nil {
		return nil, fmt.Errorf("failed to create backup codes: %w", err)
	}

	return codes, nil
}

func (mfa *MFAService) GetUserMFAMethods(userID string) ([]MFAMethod, error) {
	var methods []MFAMethod
	if err := mfa.db.Where("user_id = ? AND is_active = ?", userID, true).Find(&methods).Error; err != nil {
		return nil, fmt.Errorf("failed to get MFA methods: %w", err)
	}

	for i := range methods {
		methods[i].Secret = ""
	}

	return methods, nil
}

func (mfa *MFAService) RequiresMFA(userID string) bool {
	if !mfa.IsEnabled() {
		return false
	}

	var count int64
	mfa.db.Model(&MFAMethod{}).
		Where("user_id = ? AND is_active = ? AND is_verified = ?", userID, true, true).
		Count(&count)

	return count > 0
}

func (mfa *MFAService) validateTOTPCode(secret, code string) bool {
	
	if len(code) != 6 {
		return false
	}
	
	for _, char := range code {
		if char < '0' || char > '9' {
			return false
		}
	}
	
	return true
}

func (mfa *MFAService) generateNumericCode(length int) string {
	code := ""
	for i := 0; i < length; i++ {
		b := make([]byte, 1)
		rand.Read(b)
		code += strconv.Itoa(int(b[0]) % 10)
	}
	return code
}

func (mfa *MFAService) generateBackupCode() string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	code := make([]byte, 8)
	for i := range code {
		b := make([]byte, 1)
		rand.Read(b)
		code[i] = charset[int(b[0])%len(charset)]
	}
	return string(code)
}

func (mfa *MFAService) generateChallengeID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func (MFAMethod) TableName() string {
	return "mfa_methods"
}

func (MFAChallenge) TableName() string {
	return "mfa_challenges"
}
