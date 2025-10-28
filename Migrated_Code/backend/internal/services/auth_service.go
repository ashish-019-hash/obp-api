package services

import (
	"errors"
	"fmt"
	"time"
	"github.com/ashish-019-hash/obp-api-backend/internal/models"
	"github.com/ashish-019-hash/obp-api-backend/internal/repositories"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type AuthenticationService struct {
	db                  *gorm.DB
	authRepo            repositories.AuthRepository
	jwtSecret           string
	configService       *ConfigService
	x509Service         *X509Service
	jwksService         *JWKSService
	berlinGroupService  *BerlinGroupService
	mfaService          *MFAService
}

func NewAuthenticationService(db *gorm.DB, authRepo repositories.AuthRepository, jwtSecret string, configService *ConfigService) *AuthenticationService {
	x509Service := NewX509Service(configService)
	jwksService := NewJWKSService(configService)
	berlinGroupService := NewBerlinGroupService(db, configService)
	mfaService := NewMFAService(db, configService)
	
	return &AuthenticationService{
		db:                 db,
		authRepo:           authRepo,
		jwtSecret:          jwtSecret,
		configService:      configService,
		x509Service:        x509Service,
		jwksService:        jwksService,
		berlinGroupService: berlinGroupService,
		mfaService:         mfaService,
	}
}

func (as *AuthenticationService) ValidateJWTToken(tokenString string) (*models.User, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(as.jwtSecret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("invalid JWT token: %w", err)
	}

	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid JWT claims")
	}

	userID, ok := (*claims)["user_id"].(string)
	if !ok {
		return nil, errors.New("user_id not found in JWT claims")
	}

	var user models.User
	if err := as.db.Where("user_id = ? AND is_deleted != ?", userID, true).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("database error: %w", err)
	}

	return &user, nil
}

func (as *AuthenticationService) ValidateOAuthToken(tokenString string) (*models.User, *models.Consumer, error) {
	token, err := as.authRepo.GetTokenByValue(tokenString)
	if err != nil || token.TokenType != "access" {
		return nil, nil, errors.New("invalid OAuth token")
	}

	if time.Now().After(token.ExpiresAt) {
		return nil, nil, errors.New("token expired")
	}

	if token.User == nil {
		return nil, nil, errors.New("user not found for token")
	}

	if token.Consumer == nil {
		return nil, nil, errors.New("consumer not found for token")
	}

	if !token.Consumer.IsActive {
		return nil, nil, errors.New("consumer is inactive")
	}

	return token.User, token.Consumer, nil
}

func (as *AuthenticationService) ValidateDirectLoginToken(tokenString string) (*models.User, *models.Consumer, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(as.jwtSecret), nil
	})

	if err != nil {
		return nil, nil, fmt.Errorf("invalid DirectLogin token: %w", err)
	}

	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, nil, errors.New("invalid DirectLogin token claims")
	}

	userID, ok := (*claims)["user_id"].(string)
	if !ok {
		return nil, nil, errors.New("user_id not found in DirectLogin token")
	}

	consumerID, ok := (*claims)["consumer_id"].(string)
	if !ok {
		return nil, nil, errors.New("consumer_id not found in DirectLogin token")
	}

	var user models.User
	if err := as.db.Where("user_id = ? AND is_deleted != ?", userID, true).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, errors.New("user not found")
		}
		return nil, nil, fmt.Errorf("database error: %w", err)
	}

	var consumer models.Consumer
	if err := as.db.Where("consumer_id = ? AND is_active = ?", consumerID, true).First(&consumer).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, errors.New("consumer not found")
		}
		return nil, nil, fmt.Errorf("database error: %w", err)
	}

	return &user, &consumer, nil
}

func (as *AuthenticationService) CreateDirectLoginToken(username, password, consumerKey string) (string, error) {
	consumer, err := as.authRepo.GetConsumerByKey(consumerKey)
	if err != nil {
		return "", errors.New("invalid consumer key")
	}

	credential, err := as.authRepo.GetUserCredentialByUsername(username)
	if err != nil {
		return "", errors.New("invalid username")
	}

	if !credential.ValidatePassword(password) {
		as.RecordLoginAttempt(credential.UserID, username, "", "", "DirectLogin", false, "Invalid password")
		return "", errors.New("invalid password")
	}

	if credential.IsLocked() {
		return "", errors.New("account is locked")
	}

	tokenConfig, _ := as.configService.GetTokenConfiguration("DirectLogin")
	duration := time.Duration(tokenConfig.ExpirationSeconds) * time.Second
	claims := jwt.MapClaims{
		"user_id":     credential.UserID,
		"consumer_id": consumer.ConsumerID,
		"auth_method": "DirectLogin",
		"iat":         time.Now().Unix(),
		"exp":         time.Now().Add(duration).Unix(),
		"iss":         "OBP-API-Backend",
		"sub":         credential.UserID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(as.jwtSecret))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	dbToken := models.NewToken("DirectLogin", tokenString, consumer.ConsumerID, int64(duration.Seconds()))
	dbToken.UserID = credential.UserID

	if err := as.authRepo.CreateToken(dbToken); err != nil {
		return "", fmt.Errorf("failed to store token: %w", err)
	}

	as.RecordLoginAttempt(credential.UserID, username, "", "", "DirectLogin", true, "")

	return tokenString, nil
}

func (as *AuthenticationService) AuthenticateUser(username, password, consumerKey string) (*models.User, *models.Consumer, error) {
	var consumer models.Consumer
	if err := as.db.Where("consumer_key = ? AND is_active = ?", consumerKey, true).First(&consumer).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, errors.New("invalid consumer key")
		}
		return nil, nil, fmt.Errorf("database error: %w", err)
	}

	var userCred models.UserCredential
	if err := as.db.Where("username = ? AND is_active = ?", username, true).
		Preload("User").
		First(&userCred).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, errors.New("invalid username or password")
		}
		return nil, nil, fmt.Errorf("database error: %w", err)
	}

	if userCred.IsLocked() {
		return nil, nil, errors.New("account is temporarily locked due to failed login attempts")
	}

	if !userCred.ValidatePassword(password) {
		as.incrementFailedLoginAttempts(&userCred)
		return nil, nil, errors.New("invalid username or password")
	}

	as.resetFailedLoginAttempts(&userCred)

	now := time.Now()
	userCred.LastLoginAt = &now
	as.db.Save(&userCred)

	if userCred.User == nil {
		return nil, nil, errors.New("user not found")
	}

	if userCred.User.IsDeleted != nil && *userCred.User.IsDeleted {
		return nil, nil, errors.New("user account is deleted")
	}

	return userCred.User, &consumer, nil
}

func (as *AuthenticationService) RecordLoginAttempt(userID, username, ipAddress, userAgent, authMethod string, success bool, failureReason string) error {
	loginAttempt := models.NewLoginAttempt(userID, username, ipAddress, userAgent, authMethod, success, failureReason)
	return as.authRepo.CreateLoginAttempt(loginAttempt)
}

func (as *AuthenticationService) CreateOAuthRequestToken(consumerKey, callbackURL string) (*models.Token, error) {
	consumer, err := as.authRepo.GetConsumerByKey(consumerKey)
	if err != nil {
		return nil, errors.New("invalid consumer key")
	}

	tokenValue := as.generateSecureToken()
	tokenSecret := as.generateSecureToken()

	tokenConfig, _ := as.configService.GetTokenConfiguration("OAuth")
	token := models.NewToken("request", tokenValue, consumer.ConsumerID, tokenConfig.ExpirationSeconds)
	token.TokenSecret = tokenSecret
	token.CallbackURL = callbackURL

	if err := as.authRepo.CreateToken(token); err != nil {
		return nil, fmt.Errorf("failed to create request token: %w", err)
	}

	return token, nil
}

func (as *AuthenticationService) CreateOAuthAccessToken(oauthToken, oauthVerifier string) (*models.Token, error) {
	requestToken, err := as.authRepo.GetTokenByValue(oauthToken)
	if err != nil || requestToken.TokenType != "request" {
		return nil, errors.New("invalid request token")
	}

	if requestToken.Verifier != oauthVerifier {
		return nil, errors.New("invalid verifier")
	}

	if time.Now().After(requestToken.ExpiresAt) {
		return nil, errors.New("token expired")
	}

	requestToken.IsActive = false
	as.authRepo.UpdateToken(requestToken)

	tokenValue := as.generateSecureToken()
	tokenSecret := as.generateSecureToken()

	tokenConfig, _ := as.configService.GetTokenConfiguration("OAuth")
	accessToken := models.NewToken("access", tokenValue, requestToken.ConsumerID, tokenConfig.ExpirationSeconds)
	accessToken.TokenSecret = tokenSecret
	accessToken.UserID = requestToken.UserID

	if err := as.authRepo.CreateToken(accessToken); err != nil {
		return nil, fmt.Errorf("failed to create access token: %w", err)
	}

	return accessToken, nil
}


func (as *AuthenticationService) incrementFailedLoginAttempts(userCred *models.UserCredential) {
	userCred.FailedLoginAttempts++

	maxAttempts := as.configService.GetConfigInt("max.bad.login.attempts", 5)
	lockDurationSeconds := as.configService.GetConfigInt("user.lock.duration.seconds", 1800)
	
	if userCred.FailedLoginAttempts >= maxAttempts {
		lockUntil := time.Now().Add(time.Duration(lockDurationSeconds) * time.Second)
		userCred.LockedUntil = &lockUntil
	}

	as.db.Save(userCred)
}

func (as *AuthenticationService) resetFailedLoginAttempts(userCred *models.UserCredential) {
	userCred.FailedLoginAttempts = 0
	userCred.LockedUntil = nil
	as.db.Save(userCred)
}

func (as *AuthenticationService) generateSecureToken() string {
	return fmt.Sprintf("%d_%s", time.Now().UnixNano(), generateRandomString(32))
}

func (as *AuthenticationService) CreateConsumer(name, description, developerEmail, redirectURL, appType string) (*models.Consumer, error) {
	consumerKey := generateRandomString(32)
	consumerSecret := generateRandomString(64)
	
	consumer := models.NewConsumer(consumerKey, consumerSecret, name, developerEmail)
	consumer.Description = description
	consumer.RedirectURL = redirectURL
	consumer.AppType = appType
	
	if err := as.authRepo.CreateConsumer(consumer); err != nil {
		return nil, err
	}
	
	return consumer, nil
}

func (as *AuthenticationService) CreateUser(username, password, email, firstName, lastName, provider, providerID string) (*models.User, error) {
	user := &models.User{
		UserID:       generateSecureID(),
		Email:        email,
		FirstName:    firstName,
		LastName:     lastName,
		Provider:     provider,
		ProviderID:   providerID,
		IsActive:     true,
		ConsentGiven: false,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	
	if err := as.db.Create(user).Error; err != nil {
		return nil, err
	}
	
	bcryptCost := as.configService.GetConfigInt("bcrypt.cost", 12)
	credential, err := models.NewUserCredentialWithConfig(user.UserID, username, password, bcryptCost)
	if err != nil {
		return nil, err
	}
	
	if err := as.authRepo.CreateUserCredential(credential); err != nil {
		return nil, err
	}
	
	return user, nil
}

func (as *AuthenticationService) AuthorizeOAuthToken(oauthToken string) (string, error) {
	token, err := as.authRepo.GetTokenByValue(oauthToken)
	if err != nil || token.TokenType != "request" {
		return "", errors.New("invalid request token")
	}
	
	if time.Now().After(token.ExpiresAt) {
		return "", errors.New("token expired")
	}
	
	verifier := generateRandomString(16)
	token.Verifier = verifier
	
	if err := as.authRepo.UpdateToken(token); err != nil {
		return "", err
	}
	
	return verifier, nil
}

func (as *AuthenticationService) GetLoginAttempts(limit, offset int) ([]*models.LoginAttempt, error) {
	return as.authRepo.GetLoginAttempts(limit, offset)
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

func (as *AuthenticationService) GetEntitlementsByUserID(userID string) ([]*models.Entitlement, error) {
	return as.authRepo.GetEntitlementsByUserID(userID)
}

func (as *AuthenticationService) CheckUserEntitlement(userID, roleName string) (bool, error) {
	entitlements, err := as.GetEntitlementsByUserID(userID)
	if err != nil {
		return false, err
	}
	
	for _, entitlement := range entitlements {
		if entitlement.RoleName == roleName && entitlement.IsActive {
			return true, nil
		}
	}
	return false, nil
}

func (as *AuthenticationService) CheckConsumerScope(consumerID, roleName string) (bool, error) {
	scopes, err := as.authRepo.GetScopesByConsumerID(consumerID)
	if err != nil {
		return false, err
	}
	
	for _, scope := range scopes {
		if scope.RoleName == roleName && scope.IsActive {
			return true, nil
		}
	}
	return false, nil
}

func (as *AuthenticationService) GetTokenConfiguration(tokenType string) (*models.TokenConfiguration, error) {
	return as.configService.GetTokenConfiguration(tokenType)
}

func (as *AuthenticationService) CreateUserAuthContext(userID, consumerID, key, value string) error {
	context := models.NewUserAuthContext(userID, consumerID, key, value)
	return as.authRepo.CreateUserAuthContext(context)
}

func (as *AuthenticationService) GetUserAuthContexts(userID string) ([]*models.UserAuthContext, error) {
	return as.authRepo.GetUserAuthContexts(userID)
}

func (as *AuthenticationService) IsUserLocked(userID string) (bool, error) {
	return as.authRepo.IsUserLocked(userID)
}

func (as *AuthenticationService) LockUser(userID, lockType, reason string) error {
	lock := models.NewUserLock(userID, lockType, reason)
	return as.authRepo.CreateUserLock(lock)
}

func (as *AuthenticationService) UnlockUser(userID string) error {
	return as.authRepo.UnlockUser(userID)
}

func (as *AuthenticationService) CreateScope(consumerID, roleName string, bankID *string) error {
	scope := models.NewScope(consumerID, roleName, bankID)
	return as.authRepo.CreateScope(scope)
}

func (as *AuthenticationService) GetScopesByConsumerID(consumerID string) ([]*models.Scope, error) {
	return as.authRepo.GetScopesByConsumerID(consumerID)
}

func (as *AuthenticationService) CheckViewPermission(viewID, permissionName string) (bool, error) {
	return as.authRepo.CheckViewPermission(viewID, permissionName)
}

func (as *AuthenticationService) CreateViewPermission(viewID, permissionName string, bankID, accountID *string) error {
	permission := models.NewViewPermission(viewID, permissionName, bankID, accountID)
	return as.authRepo.CreateViewPermission(permission)
}

func (as *AuthenticationService) ValidateAuthenticationTypeForOperation(operationID, authType string) (bool, error) {
	validation, err := as.authRepo.GetAuthTypeValidationByOperation(operationID)
	if err != nil {
		return true, nil
	}
	
	allowedTypes := validation.GetAuthTypes()
	for _, allowedType := range allowedTypes {
		if allowedType == authType {
			return true, nil
		}
	}
	return false, nil
}


func (as *AuthenticationService) ValidateOIDCToken(tokenString string) (*models.User, *models.Consumer, error) {
	token, claims, err := as.jwksService.ValidateOIDCToken(tokenString)
	if err != nil {
		return nil, nil, fmt.Errorf("OIDC token validation failed: %w", err)
	}

	sub, ok := claims["sub"].(string)
	if !ok {
		return nil, nil, errors.New("missing subject in OIDC token")
	}

	email, _ := claims["email"].(string)
	name, _ := claims["name"].(string)

	var user models.User
	err = as.db.Where("provider_id = ? AND provider = ?", sub, "oidc").First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		user = models.User{
			UserID:     fmt.Sprintf("oidc_%s", sub),
			Email:      email,
			FirstName:  name,
			Provider:   "oidc",
			ProviderID: sub,
			IsActive:   true,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}
		if err := as.db.Create(&user).Error; err != nil {
			return nil, nil, fmt.Errorf("failed to create OIDC user: %w", err)
		}
	} else if err != nil {
		return nil, nil, fmt.Errorf("database error: %w", err)
	}

	clientID, _ := claims["aud"].(string)
	var consumer models.Consumer
	err = as.db.Where("consumer_key = ?", clientID).First(&consumer).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		consumer = models.Consumer{
			ConsumerID:  fmt.Sprintf("oidc_%s", clientID),
			ConsumerKey: clientID,
			Name:        "OIDC Client",
			AppType:     "Web",
			IsActive:    true,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		if err := as.db.Create(&consumer).Error; err != nil {
			return nil, nil, fmt.Errorf("failed to create OIDC consumer: %w", err)
		}
	}

	return &user, &consumer, nil
}

func (as *AuthenticationService) ValidateCertificateAuth(certPEM string) (*models.User, *models.Consumer, error) {
	certInfo, err := as.x509Service.ValidateCertificate(certPEM)
	if err != nil {
		return nil, nil, fmt.Errorf("certificate validation failed: %w", err)
	}

	if !certInfo.IsValid {
		return nil, nil, fmt.Errorf("invalid certificate: %s", certInfo.ValidationError)
	}

	var consumer models.Consumer
	err = as.db.Where("client_certificate = ?", certPEM).First(&consumer).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil, errors.New("no consumer found for certificate")
	} else if err != nil {
		return nil, nil, fmt.Errorf("database error: %w", err)
	}

	var user models.User
	err = as.db.Where("email = ?", certInfo.Email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		user = models.User{
			UserID:    fmt.Sprintf("cert_%s", certInfo.SerialNumber),
			Email:     certInfo.Email,
			FirstName: certInfo.CommonName,
			Provider:  "certificate",
			ProviderID: certInfo.SerialNumber,
			IsActive:  true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		if err := as.db.Create(&user).Error; err != nil {
			return nil, nil, fmt.Errorf("failed to create certificate user: %w", err)
		}
	}

	return &user, &consumer, nil
}

func (as *AuthenticationService) CreateBerlinGroupConsent(userID, consumerID string, access interface{}) (string, error) {
	consent := &models.Consent{
		ConsentID:   fmt.Sprintf("bg_%d", time.Now().Unix()),
		UserID:      userID,
		ConsumerID:  consumerID,
		Status:      "INITIATED",
		ConsentType: "BERLIN_GROUP",
		ValidFrom:   time.Now(),
		ValidUntil:  time.Now().Add(24 * time.Hour),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := as.db.Create(consent).Error; err != nil {
		return "", fmt.Errorf("failed to create Berlin Group consent: %w", err)
	}

	return consent.ConsentID, nil
}

func (as *AuthenticationService) RequiresMFA(userID string) bool {
	return as.mfaService.RequiresMFA(userID)
}

func (as *AuthenticationService) SetupMFA(userID, method string) (interface{}, error) {
	switch method {
	case "TOTP":
		secret, qrURL, err := as.mfaService.SetupTOTP(userID)
		if err != nil {
			return nil, err
		}
		return map[string]string{
			"secret": secret,
			"qr_url": qrURL,
		}, nil
	case "SMS":
		return nil, errors.New("SMS MFA setup requires phone number")
	default:
		return nil, errors.New("unsupported MFA method")
	}
}

func (as *AuthenticationService) VerifyMFA(userID, method, code string) error {
	switch method {
	case "TOTP":
		return as.mfaService.VerifyTOTP(userID, code)
	case "SMS":
		return errors.New("SMS MFA verification requires challenge ID")
	default:
		return errors.New("unsupported MFA method")
	}
}
