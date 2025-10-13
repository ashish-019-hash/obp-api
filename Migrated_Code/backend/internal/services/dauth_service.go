package services

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/ashish-019-hash/obp-api-backend/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type DAuthService struct {
	db            *gorm.DB
	configService *ConfigService
	randomService *SecureRandomService
	jwtSecret     string
}

type DAuthPayload struct {
	UserID     string `json:"user_id"`
	ConsumerID string `json:"consumer_id"`
	AppName    string `json:"app_name"`
	Timestamp  int64  `json:"timestamp"`
	Nonce      string `json:"nonce"`
}

type DynamicConsumerRequest struct {
	AppName     string `json:"app_name" binding:"required"`
	AppType     string `json:"app_type"`
	Description string `json:"description"`
	RedirectURL string `json:"redirect_url"`
}

func NewDAuthService(db *gorm.DB, configService *ConfigService, randomService *SecureRandomService, jwtSecret string) *DAuthService {
	return &DAuthService{
		db:            db,
		configService: configService,
		randomService: randomService,
		jwtSecret:     jwtSecret,
	}
}

func (ds *DAuthService) CreateDynamicConsumer(req DynamicConsumerRequest, userID string) (*models.Consumer, error) {
	consumerKey, err := ds.randomService.ConsumerKey()
	if err != nil {
		return nil, fmt.Errorf("failed to generate consumer key: %w", err)
	}
	
	consumerSecret, err := ds.randomService.ConsumerSecret()
	if err != nil {
		return nil, fmt.Errorf("failed to generate consumer secret: %w", err)
	}
	
	consumer := models.NewConsumer(consumerKey, consumerSecret, req.AppName, "")
	consumer.Description = req.Description
	consumer.RedirectURL = req.RedirectURL
	consumer.AppType = req.AppType
	consumer.CreatedByUserID = userID
	consumer.IsActive = true
	
	if err := ds.db.Create(consumer).Error; err != nil {
		return nil, fmt.Errorf("failed to create dynamic consumer: %w", err)
	}
	
	return consumer, nil
}

func (ds *DAuthService) CreateDAuthToken(userID, consumerID, appName string) (string, error) {
	nonce, err := ds.randomService.Nonce()
	if err != nil {
		return "", fmt.Errorf("failed to generate nonce: %w", err)
	}
	
	payload := DAuthPayload{
		UserID:     userID,
		ConsumerID: consumerID,
		AppName:    appName,
		Timestamp:  time.Now().Unix(),
		Nonce:      nonce,
	}
	
	expirationSeconds := ds.configService.GetConfigInt("dauth.token.expiration.seconds", 3600)
	
	claims := jwt.MapClaims{
		"user_id":     payload.UserID,
		"consumer_id": payload.ConsumerID,
		"app_name":    payload.AppName,
		"timestamp":   payload.Timestamp,
		"nonce":       payload.Nonce,
		"iat":         time.Now().Unix(),
		"exp":         time.Now().Add(time.Duration(expirationSeconds) * time.Second).Unix(),
		"iss":         "obp-api",
		"aud":         "obp-api-client",
		"auth_type":   "DAuth",
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(ds.jwtSecret))
	if err != nil {
		return "", fmt.Errorf("failed to sign DAuth token: %w", err)
	}
	
	tokenModel := models.NewToken("dauth", tokenString, consumerID, int64(expirationSeconds))
	tokenModel.UserID = userID
	
	if err := ds.db.Create(tokenModel).Error; err != nil {
		return "", fmt.Errorf("failed to store DAuth token: %w", err)
	}
	
	return tokenString, nil
}

func (ds *DAuthService) ValidateDAuthToken(tokenString string) (*models.User, *models.Consumer, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(ds.jwtSecret), nil
	})
	
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse DAuth token: %w", err)
	}
	
	if !token.Valid {
		return nil, nil, fmt.Errorf("invalid DAuth token")
	}
	
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, nil, fmt.Errorf("invalid DAuth token claims")
	}
	
	authType, ok := claims["auth_type"].(string)
	if !ok || authType != "DAuth" {
		return nil, nil, fmt.Errorf("invalid auth_type in DAuth token")
	}
	
	userID, ok := claims["user_id"].(string)
	if !ok {
		return nil, nil, fmt.Errorf("missing user_id in DAuth token")
	}
	
	consumerID, ok := claims["consumer_id"].(string)
	if !ok {
		return nil, nil, fmt.Errorf("missing consumer_id in DAuth token")
	}
	
	var tokenModel models.Token
	if err := ds.db.Where("token_value = ? AND token_type = ? AND is_active = ?", 
		tokenString, "dauth", true).First(&tokenModel).Error; err != nil {
		return nil, nil, fmt.Errorf("DAuth token not found or inactive: %w", err)
	}
	
	if tokenModel.ExpiresAt.Before(time.Now()) {
		return nil, nil, fmt.Errorf("DAuth token expired")
	}
	
	var user models.User
	if err := ds.db.Where("user_id = ? AND is_active = ?", userID, true).First(&user).Error; err != nil {
		return nil, nil, fmt.Errorf("user not found or inactive: %w", err)
	}
	
	var consumer models.Consumer
	if err := ds.db.Where("consumer_id = ? AND is_active = ?", consumerID, true).First(&consumer).Error; err != nil {
		return nil, nil, fmt.Errorf("consumer not found or inactive: %w", err)
	}
	
	return &user, &consumer, nil
}

func (ds *DAuthService) RevokeDAuthToken(tokenString string) error {
	return ds.db.Model(&models.Token{}).
		Where("token_value = ? AND token_type = ?", tokenString, "dauth").
		Update("is_active", false).Error
}

func (ds *DAuthService) GetDAuthTokensByUser(userID string) ([]models.Token, error) {
	var tokens []models.Token
	err := ds.db.Where("user_id = ? AND token_type = ? AND is_active = ?", 
		userID, "dauth", true).Find(&tokens).Error
	
	return tokens, err
}

func (ds *DAuthService) GetDAuthTokensByConsumer(consumerID string) ([]models.Token, error) {
	var tokens []models.Token
	err := ds.db.Where("consumer_id = ? AND token_type = ? AND is_active = ?", 
		consumerID, "dauth", true).Find(&tokens).Error
	
	return tokens, err
}

func (ds *DAuthService) CleanupExpiredDAuthTokens() error {
	return ds.db.Where("token_type = ? AND expires_at < ?", "dauth", time.Now()).
		Delete(&models.Token{}).Error
}

func (ds *DAuthService) ValidateDAuthPayload(payloadJSON string) (*DAuthPayload, error) {
	var payload DAuthPayload
	if err := json.Unmarshal([]byte(payloadJSON), &payload); err != nil {
		return nil, fmt.Errorf("invalid DAuth payload JSON: %w", err)
	}
	
	if payload.UserID == "" {
		return nil, fmt.Errorf("missing user_id in DAuth payload")
	}
	
	if payload.ConsumerID == "" {
		return nil, fmt.Errorf("missing consumer_id in DAuth payload")
	}
	
	if payload.AppName == "" {
		return nil, fmt.Errorf("missing app_name in DAuth payload")
	}
	
	if payload.Nonce == "" {
		return nil, fmt.Errorf("missing nonce in DAuth payload")
	}
	
	now := time.Now().Unix()
	if payload.Timestamp < now-300 || payload.Timestamp > now+300 {
		return nil, fmt.Errorf("DAuth payload timestamp out of acceptable range")
	}
	
	return &payload, nil
}

func (ds *DAuthService) IsDAuthEnabled() bool {
	return ds.configService.GetConfigBool("dauth.enabled", false)
}
