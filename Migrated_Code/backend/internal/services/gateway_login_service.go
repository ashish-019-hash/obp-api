package services

import (
	"fmt"
	"time"

	"github.com/ashish-019-hash/obp-api-backend/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type GatewayLoginService struct {
	db            *gorm.DB
	configService *ConfigService
	randomService *SecureRandomService
	jwtSecret     string
}

type GatewayLoginPayload struct {
	UserID       string `json:"user_id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	CBSToken     string `json:"cbs_token"`
	BankID       string `json:"bank_id"`
	IsFirst      bool   `json:"is_first"`
	Timestamp    int64  `json:"timestamp"`
}

func NewGatewayLoginService(db *gorm.DB, configService *ConfigService, randomService *SecureRandomService, jwtSecret string) *GatewayLoginService {
	return &GatewayLoginService{
		db:            db,
		configService: configService,
		randomService: randomService,
		jwtSecret:     jwtSecret,
	}
}

func (gls *GatewayLoginService) CreateGatewayJWT(payload GatewayLoginPayload) (string, error) {
	if !gls.IsGatewayLoginEnabled() {
		return "", fmt.Errorf("gateway login is disabled")
	}
	
	expirationSeconds := gls.configService.GetConfigInt("gateway.token.expiration.seconds", 3600)
	
	claims := jwt.MapClaims{
		"user_id":    payload.UserID,
		"username":   payload.Username,
		"email":      payload.Email,
		"first_name": payload.FirstName,
		"last_name":  payload.LastName,
		"cbs_token":  payload.CBSToken,
		"bank_id":    payload.BankID,
		"is_first":   payload.IsFirst,
		"timestamp":  payload.Timestamp,
		"iat":        time.Now().Unix(),
		"exp":        time.Now().Add(time.Duration(expirationSeconds) * time.Second).Unix(),
		"iss":        "obp-gateway",
		"aud":        "obp-api",
		"auth_type":  "GatewayLogin",
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(gls.jwtSecret))
	if err != nil {
		return "", fmt.Errorf("failed to sign gateway JWT: %w", err)
	}
	
	return tokenString, nil
}

func (gls *GatewayLoginService) ValidateGatewayToken(tokenString string) (*models.User, *models.Consumer, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(gls.jwtSecret), nil
	})
	
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse gateway token: %w", err)
	}
	
	if !token.Valid {
		return nil, nil, fmt.Errorf("invalid gateway token")
	}
	
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, nil, fmt.Errorf("invalid gateway token claims")
	}
	
	authType, ok := claims["auth_type"].(string)
	if !ok || authType != "GatewayLogin" {
		return nil, nil, fmt.Errorf("invalid auth_type in gateway token")
	}
	
	userID, ok := claims["user_id"].(string)
	if !ok {
		return nil, nil, fmt.Errorf("missing user_id in gateway token")
	}
	
	username, ok := claims["username"].(string)
	if !ok {
		return nil, nil, fmt.Errorf("missing username in gateway token")
	}
	
	email, ok := claims["email"].(string)
	if !ok {
		return nil, nil, fmt.Errorf("missing email in gateway token")
	}
	
	isFirst, _ := claims["is_first"].(bool)
	
	user, err := gls.getOrCreateGatewayUser(userID, username, email, claims, isFirst)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get/create gateway user: %w", err)
	}
	
	consumer, err := gls.getOrCreateGatewayConsumer(userID)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get/create gateway consumer: %w", err)
	}
	
	return user, consumer, nil
}

func (gls *GatewayLoginService) getOrCreateGatewayUser(userID, username, email string, claims jwt.MapClaims, isFirst bool) (*models.User, error) {
	var user models.User
	err := gls.db.Where("user_id = ?", userID).First(&user).Error
	
	if err == gorm.ErrRecordNotFound || isFirst {
		firstName, _ := claims["first_name"].(string)
		lastName, _ := claims["last_name"].(string)
		bankID, _ := claims["bank_id"].(string)
		
		user = models.User{
			UserID:       userID,
			Email:        email,
			FirstName:    firstName,
			LastName:     lastName,
			Provider:     "gateway",
			ProviderID:   username,
			IsActive:     true,
			ConsentGiven: true,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}
		
		if bankID != "" {
			user.ProviderID = fmt.Sprintf("%s@%s", username, bankID)
		}
		
		if err := gls.db.Create(&user).Error; err != nil {
			return nil, fmt.Errorf("failed to create gateway user: %w", err)
		}
	} else if err != nil {
		return nil, fmt.Errorf("failed to query gateway user: %w", err)
	} else {
		firstName, _ := claims["first_name"].(string)
		lastName, _ := claims["last_name"].(string)
		
		if firstName != "" {
			user.FirstName = firstName
		}
		if lastName != "" {
			user.LastName = lastName
		}
		
		user.UpdatedAt = time.Now()
		
		if err := gls.db.Save(&user).Error; err != nil {
			return nil, fmt.Errorf("failed to update gateway user: %w", err)
		}
	}
	
	return &user, nil
}

func (gls *GatewayLoginService) getOrCreateGatewayConsumer(userID string) (*models.Consumer, error) {
	var consumer models.Consumer
	consumerID := fmt.Sprintf("gateway_%s", userID)
	
	err := gls.db.Where("consumer_id = ?", consumerID).First(&consumer).Error
	
	if err == gorm.ErrRecordNotFound {
		consumerKey, err := gls.randomService.ConsumerKey()
		if err != nil {
			return nil, fmt.Errorf("failed to generate consumer key: %w", err)
		}
		
		consumerSecret, err := gls.randomService.ConsumerSecret()
		if err != nil {
			return nil, fmt.Errorf("failed to generate consumer secret: %w", err)
		}
		
		consumer = *models.NewConsumer(consumerKey, consumerSecret, "Gateway Consumer", "")
		consumer.ConsumerID = consumerID
		consumer.Description = "Auto-generated consumer for Gateway Login"
		consumer.AppType = "Gateway"
		consumer.CreatedByUserID = userID
		consumer.IsActive = true
		
		if err := gls.db.Create(&consumer).Error; err != nil {
			return nil, fmt.Errorf("failed to create gateway consumer: %w", err)
		}
	} else if err != nil {
		return nil, fmt.Errorf("failed to query gateway consumer: %w", err)
	}
	
	return &consumer, nil
}

func (gls *GatewayLoginService) ValidateGatewayHost(host string) bool {
	whitelist := gls.configService.GetConfig("gateway.host.whitelist", "localhost")
	
	if whitelist == "*" {
		return true
	}
	
	for _, allowedHost := range []string{whitelist} {
		if host == allowedHost {
			return true
		}
	}
	
	return false
}

func (gls *GatewayLoginService) IsGatewayLoginEnabled() bool {
	return gls.configService.GetConfigBool("gateway.login.enabled", false)
}

func (gls *GatewayLoginService) CreateGatewaySession(userID, consumerID, ipAddress, userAgent string) (*models.APISession, error) {
	timeoutMinutes := gls.configService.GetConfigInt("gateway.session.timeout.minutes", 60)
	
	session := models.NewAPISession(userID, consumerID, ipAddress, userAgent, timeoutMinutes)
	
	if err := gls.db.Create(session).Error; err != nil {
		return nil, fmt.Errorf("failed to create gateway session: %w", err)
	}
	
	return session, nil
}

func (gls *GatewayLoginService) RefreshGatewayAccounts(userID, cbsToken string) error {
	
	
	return nil
}

func (gls *GatewayLoginService) GetGatewayTokensByUser(userID string) ([]models.Token, error) {
	var tokens []models.Token
	err := gls.db.Where("user_id = ? AND token_type = ? AND is_active = ?", 
		userID, "gateway", true).Find(&tokens).Error
	
	return tokens, err
}

func (gls *GatewayLoginService) RevokeGatewayToken(tokenString string) error {
	return gls.db.Model(&models.Token{}).
		Where("token_value = ? AND token_type = ?", tokenString, "gateway").
		Update("is_active", false).Error
}
