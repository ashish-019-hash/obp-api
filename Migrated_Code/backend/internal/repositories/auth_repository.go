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

	CreateScope(scope *models.Scope) error
	GetScopesByConsumerID(consumerID string) ([]*models.Scope, error)
	GetScopeByID(scopeID string) (*models.Scope, error)
	DeleteScope(scopeID string) error

	CreateViewPermission(permission *models.ViewPermission) error
	GetViewPermissionsByViewID(viewID string) ([]*models.ViewPermission, error)
	CheckViewPermission(viewID, permissionName string) (bool, error)
	DeleteViewPermission(permissionID string) error

	CreateUserAuthContext(context *models.UserAuthContext) error
	GetUserAuthContexts(userID string) ([]*models.UserAuthContext, error)
	DeleteUserAuthContexts(userID string) error
	DeleteUserAuthContextByID(contextID string) error

	CreateConsentAuthContext(context *models.ConsentAuthContext) error
	GetConsentAuthContexts(consentID string) ([]*models.ConsentAuthContext, error)
	DeleteConsentAuthContexts(consentID string) error

	CreateUserLock(lock *models.UserLock) error
	GetUserLocksByUserID(userID string) ([]*models.UserLock, error)
	IsUserLocked(userID string) (bool, error)
	UnlockUser(userID string) error

	CreateAuthTypeValidation(validation *models.AuthenticationTypeValidation) error
	GetAuthTypeValidationByOperation(operationID string) (*models.AuthenticationTypeValidation, error)
	UpdateAuthTypeValidation(validation *models.AuthenticationTypeValidation) error
	DeleteAuthTypeValidation(operationID string) error
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) CreateScope(scope *models.Scope) error {
	return r.db.Create(scope).Error
}

func (r *authRepository) GetScopesByConsumerID(consumerID string) ([]*models.Scope, error) {
	var scopes []*models.Scope
	err := r.db.Where("consumer_id = ? AND is_active = ?", consumerID, true).Find(&scopes).Error
	return scopes, err
}

func (r *authRepository) GetScopeByID(scopeID string) (*models.Scope, error) {
	var scope models.Scope
	err := r.db.Where("scope_id = ?", scopeID).First(&scope).Error
	if err != nil {
		return nil, err
	}
	return &scope, nil
}

func (r *authRepository) DeleteScope(scopeID string) error {
	return r.db.Where("scope_id = ?", scopeID).Delete(&models.Scope{}).Error
}

func (r *authRepository) CreateViewPermission(permission *models.ViewPermission) error {
	return r.db.Create(permission).Error
}

func (r *authRepository) GetViewPermissionsByViewID(viewID string) ([]*models.ViewPermission, error) {
	var permissions []*models.ViewPermission
	err := r.db.Where("view_id = ? AND is_active = ?", viewID, true).Find(&permissions).Error
	return permissions, err
}

func (r *authRepository) CheckViewPermission(viewID, permissionName string) (bool, error) {
	var count int64
	err := r.db.Model(&models.ViewPermission{}).Where("view_id = ? AND permission_name = ? AND is_active = ?", viewID, permissionName, true).Count(&count).Error
	return count > 0, err
}

func (r *authRepository) DeleteViewPermission(permissionID string) error {
	return r.db.Where("permission_id = ?", permissionID).Delete(&models.ViewPermission{}).Error
}

func (r *authRepository) CreateUserAuthContext(context *models.UserAuthContext) error {
	return r.db.Create(context).Error
}

func (r *authRepository) GetUserAuthContexts(userID string) ([]*models.UserAuthContext, error) {
	var contexts []*models.UserAuthContext
	err := r.db.Where("user_id = ?", userID).Find(&contexts).Error
	return contexts, err
}

func (r *authRepository) DeleteUserAuthContexts(userID string) error {
	return r.db.Where("user_id = ?", userID).Delete(&models.UserAuthContext{}).Error
}

func (r *authRepository) DeleteUserAuthContextByID(contextID string) error {
	return r.db.Where("context_id = ?", contextID).Delete(&models.UserAuthContext{}).Error
}

func (r *authRepository) CreateConsentAuthContext(context *models.ConsentAuthContext) error {
	return r.db.Create(context).Error
}

func (r *authRepository) GetConsentAuthContexts(consentID string) ([]*models.ConsentAuthContext, error) {
	var contexts []*models.ConsentAuthContext
	err := r.db.Where("consent_id = ?", consentID).Find(&contexts).Error
	return contexts, err
}

func (r *authRepository) DeleteConsentAuthContexts(consentID string) error {
	return r.db.Where("consent_id = ?", consentID).Delete(&models.ConsentAuthContext{}).Error
}

func (r *authRepository) CreateUserLock(lock *models.UserLock) error {
	return r.db.Create(lock).Error
}

func (r *authRepository) GetUserLocksByUserID(userID string) ([]*models.UserLock, error) {
	var locks []*models.UserLock
	err := r.db.Where("user_id = ? AND is_active = ?", userID, true).Find(&locks).Error
	return locks, err
}

func (r *authRepository) IsUserLocked(userID string) (bool, error) {
	var count int64
	err := r.db.Model(&models.UserLock{}).Where("user_id = ? AND is_active = ?", userID, true).Count(&count).Error
	return count > 0, err
}

func (r *authRepository) UnlockUser(userID string) error {
	return r.db.Model(&models.UserLock{}).Where("user_id = ?", userID).Update("is_active", false).Error
}

func (r *authRepository) CreateAuthTypeValidation(validation *models.AuthenticationTypeValidation) error {
	return r.db.Create(validation).Error
}

func (r *authRepository) GetAuthTypeValidationByOperation(operationID string) (*models.AuthenticationTypeValidation, error) {
	var validation models.AuthenticationTypeValidation
	err := r.db.Where("operation_id = ? AND is_active = ?", operationID, true).First(&validation).Error
	if err != nil {
		return nil, err
	}
	return &validation, nil
}

func (r *authRepository) UpdateAuthTypeValidation(validation *models.AuthenticationTypeValidation) error {
	return r.db.Save(validation).Error
}

func (r *authRepository) DeleteAuthTypeValidation(operationID string) error {
	return r.db.Where("operation_id = ?", operationID).Delete(&models.AuthenticationTypeValidation{}).Error
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
