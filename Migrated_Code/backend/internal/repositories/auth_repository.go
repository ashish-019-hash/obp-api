package repositories

import (
	"github.com/ashish-019-hash/obp-api-backend/internal/models"
	"gorm.io/gorm"
)

type AuthRepository interface {
	CreateConsumer(consumer *models.Consumer) error
	GetConsumerByKey(consumerKey string) (*models.Consumer, error)
	GetConsumerByID(consumerID string) (*models.Consumer, error)
	UpdateConsumer(consumer *models.Consumer) error

	CreateToken(token *models.Token) error
	GetTokenByValue(tokenValue string) (*models.Token, error)
	GetTokensByConsumerID(consumerID string) ([]*models.Token, error)
	UpdateToken(token *models.Token) error
	DeleteToken(tokenValue string) error

	CreateUserCredential(credential *models.UserCredential) error
	GetUserCredentialByUsername(username string) (*models.UserCredential, error)
	GetUserCredentialByUserID(userID string) (*models.UserCredential, error)
	UpdateUserCredential(credential *models.UserCredential) error

	CreateEntitlement(entitlement *models.Entitlement) error
	GetEntitlementsByUserID(userID string) ([]*models.Entitlement, error)
	GetEntitlementByID(entitlementID string) (*models.Entitlement, error)
	UpdateEntitlement(entitlement *models.Entitlement) error
	DeleteEntitlement(entitlementID string) error

	CreateLoginAttempt(attempt *models.LoginAttempt) error
	GetLoginAttemptsByUserID(userID string, limit, offset int) ([]*models.LoginAttempt, error)
	GetLoginAttempts(limit, offset int) ([]*models.LoginAttempt, error)
	GetFailedLoginAttempts(userID string, since int64) (int, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) CreateConsumer(consumer *models.Consumer) error {
	return r.db.Create(consumer).Error
}

func (r *authRepository) GetConsumerByKey(consumerKey string) (*models.Consumer, error) {
	var consumer models.Consumer
	err := r.db.Where("consumer_key = ? AND is_active = ?", consumerKey, true).First(&consumer).Error
	if err != nil {
		return nil, err
	}
	return &consumer, nil
}

func (r *authRepository) GetConsumerByID(consumerID string) (*models.Consumer, error) {
	var consumer models.Consumer
	err := r.db.Where("consumer_id = ?", consumerID).First(&consumer).Error
	if err != nil {
		return nil, err
	}
	return &consumer, nil
}

func (r *authRepository) UpdateConsumer(consumer *models.Consumer) error {
	return r.db.Save(consumer).Error
}

func (r *authRepository) CreateToken(token *models.Token) error {
	return r.db.Create(token).Error
}

func (r *authRepository) GetTokenByValue(tokenValue string) (*models.Token, error) {
	var token models.Token
	err := r.db.Preload("Consumer").Preload("User").Where("token_value = ? AND is_active = ?", tokenValue, true).First(&token).Error
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (r *authRepository) GetTokensByConsumerID(consumerID string) ([]*models.Token, error) {
	var tokens []*models.Token
	err := r.db.Where("consumer_id = ? AND is_active = ?", consumerID, true).Find(&tokens).Error
	return tokens, err
}

func (r *authRepository) UpdateToken(token *models.Token) error {
	return r.db.Save(token).Error
}

func (r *authRepository) DeleteToken(tokenValue string) error {
	return r.db.Model(&models.Token{}).Where("token_value = ?", tokenValue).Update("is_active", false).Error
}

func (r *authRepository) CreateUserCredential(credential *models.UserCredential) error {
	return r.db.Create(credential).Error
}

func (r *authRepository) GetUserCredentialByUsername(username string) (*models.UserCredential, error) {
	var credential models.UserCredential
	err := r.db.Preload("User").Where("username = ? AND is_active = ?", username, true).First(&credential).Error
	if err != nil {
		return nil, err
	}
	return &credential, nil
}

func (r *authRepository) GetUserCredentialByUserID(userID string) (*models.UserCredential, error) {
	var credential models.UserCredential
	err := r.db.Preload("User").Where("user_id = ? AND is_active = ?", userID, true).First(&credential).Error
	if err != nil {
		return nil, err
	}
	return &credential, nil
}

func (r *authRepository) UpdateUserCredential(credential *models.UserCredential) error {
	return r.db.Save(credential).Error
}

func (r *authRepository) CreateEntitlement(entitlement *models.Entitlement) error {
	return r.db.Create(entitlement).Error
}

func (r *authRepository) GetEntitlementsByUserID(userID string) ([]*models.Entitlement, error) {
	var entitlements []*models.Entitlement
	err := r.db.Where("user_id = ? AND is_active = ?", userID, true).Find(&entitlements).Error
	return entitlements, err
}

func (r *authRepository) GetEntitlementByID(entitlementID string) (*models.Entitlement, error) {
	var entitlement models.Entitlement
	err := r.db.Where("entitlement_id = ?", entitlementID).First(&entitlement).Error
	if err != nil {
		return nil, err
	}
	return &entitlement, nil
}

func (r *authRepository) UpdateEntitlement(entitlement *models.Entitlement) error {
	return r.db.Save(entitlement).Error
}

func (r *authRepository) DeleteEntitlement(entitlementID string) error {
	return r.db.Model(&models.Entitlement{}).Where("entitlement_id = ?", entitlementID).Update("is_active", false).Error
}

func (r *authRepository) CreateLoginAttempt(attempt *models.LoginAttempt) error {
	return r.db.Create(attempt).Error
}

func (r *authRepository) GetLoginAttemptsByUserID(userID string, limit, offset int) ([]*models.LoginAttempt, error) {
	var attempts []*models.LoginAttempt
	err := r.db.Where("user_id = ?", userID).Order("attempted_at DESC").Limit(limit).Offset(offset).Find(&attempts).Error
	return attempts, err
}

func (r *authRepository) GetLoginAttempts(limit, offset int) ([]*models.LoginAttempt, error) {
	var attempts []*models.LoginAttempt
	err := r.db.Order("attempted_at DESC").Limit(limit).Offset(offset).Find(&attempts).Error
	return attempts, err
}

func (r *authRepository) GetFailedLoginAttempts(userID string, since int64) (int, error) {
	var count int64
	err := r.db.Model(&models.LoginAttempt{}).Where("user_id = ? AND success = ? AND attempted_at > ?", userID, false, since).Count(&count).Error
	return int(count), err
}
