package services

import (
	"strconv"
	"time"

	"github.com/ashish-019-hash/obp-api-backend/internal/models"
	"gorm.io/gorm"
)

type ConfigService struct {
	db *gorm.DB
}

func NewConfigService(db *gorm.DB) *ConfigService {
	return &ConfigService{db: db}
}

func (cs *ConfigService) GetConfig(key, defaultValue string) string {
	var config models.AuthenticationConfig
	if err := cs.db.Where("config_key = ? AND is_active = ?", key, true).First(&config).Error; err != nil {
		return defaultValue
	}
	return config.ConfigValue
}

func (cs *ConfigService) GetConfigInt(key string, defaultValue int) int {
	value := cs.GetConfig(key, strconv.Itoa(defaultValue))
	if intValue, err := strconv.Atoi(value); err == nil {
		return intValue
	}
	return defaultValue
}

func (cs *ConfigService) GetConfigBool(key string, defaultValue bool) bool {
	value := cs.GetConfig(key, strconv.FormatBool(defaultValue))
	if boolValue, err := strconv.ParseBool(value); err == nil {
		return boolValue
	}
	return defaultValue
}

func (cs *ConfigService) GetConfigDuration(key string, defaultValue time.Duration) time.Duration {
	value := cs.GetConfig(key, defaultValue.String())
	if duration, err := time.ParseDuration(value); err == nil {
		return duration
	}
	return defaultValue
}

func (cs *ConfigService) SetConfig(key, value, configType, description string) error {
	var config models.AuthenticationConfig
	err := cs.db.Where("config_key = ?", key).First(&config).Error
	
	if err == gorm.ErrRecordNotFound {
		config = *models.NewAuthenticationConfig(key, value, configType, description)
		return cs.db.Create(&config).Error
	} else if err != nil {
		return err
	}
	
	config.ConfigValue = value
	config.UpdatedAt = time.Now()
	return cs.db.Save(&config).Error
}

func (cs *ConfigService) GetConsumerRateLimit(consumerID string) (*models.ConsumerRateLimit, error) {
	var rateLimit models.ConsumerRateLimit
	err := cs.db.Where("consumer_id = ? AND is_active = ?", consumerID, true).First(&rateLimit).Error
	if err == gorm.ErrRecordNotFound {
		return models.NewConsumerRateLimit(consumerID, 100, 1000, 10000), nil
	}
	return &rateLimit, err
}

func (cs *ConfigService) SetConsumerRateLimit(consumerID string, perMinute, perHour, perDay int) error {
	var rateLimit models.ConsumerRateLimit
	err := cs.db.Where("consumer_id = ?", consumerID).First(&rateLimit).Error
	
	if err == gorm.ErrRecordNotFound {
		rateLimit = *models.NewConsumerRateLimit(consumerID, perMinute, perHour, perDay)
		return cs.db.Create(&rateLimit).Error
	} else if err != nil {
		return err
	}
	
	rateLimit.RequestsPerMinute = perMinute
	rateLimit.RequestsPerHour = perHour
	rateLimit.RequestsPerDay = perDay
	rateLimit.UpdatedAt = time.Now()
	return cs.db.Save(&rateLimit).Error
}

func (cs *ConfigService) GetSecuritySetting(key, defaultValue string) string {
	var setting models.SecuritySettings
	if err := cs.db.Where("setting_key = ? AND is_active = ?", key, true).First(&setting).Error; err != nil {
		return defaultValue
	}
	return setting.SettingValue
}

func (cs *ConfigService) GetSecuritySettingInt(key string, defaultValue int) int {
	value := cs.GetSecuritySetting(key, strconv.Itoa(defaultValue))
	if intValue, err := strconv.Atoi(value); err == nil {
		return intValue
	}
	return defaultValue
}

func (cs *ConfigService) GetTokenConfiguration(tokenType string) (*models.TokenConfiguration, error) {
	var tokenConfig models.TokenConfiguration
	err := cs.db.Where("token_type = ? AND is_active = ?", tokenType, true).First(&tokenConfig).Error
	if err == gorm.ErrRecordNotFound {
		switch tokenType {
		case "DirectLogin":
			return models.NewTokenConfiguration(tokenType, 2419200, false, 0), nil
		case "OAuth":
			return models.NewTokenConfiguration(tokenType, 3600, true, 86400), nil
		case "JWT":
			return models.NewTokenConfiguration(tokenType, 86400, false, 0), nil
		default:
			return models.NewTokenConfiguration(tokenType, 3600, false, 0), nil
		}
	}
	return &tokenConfig, err
}

func (cs *ConfigService) InitializeDefaultConfigs() error {
	defaults := []struct {
		key, value, configType, description string
	}{
		{"max.bad.login.attempts", "5", "int", "Maximum failed login attempts before lockout"},
		{"user.lock.duration.seconds", "1800", "int", "User lock duration in seconds (30 minutes)"},
		{"bcrypt.cost", "12", "int", "Bcrypt hashing cost"},
		{"direct.login.token.expiration.seconds", "2419200", "int", "DirectLogin token expiration (4 weeks)"},
		{"oauth.token.expiration.seconds", "3600", "int", "OAuth token expiration (1 hour)"},
		{"jwt.token.expiration.seconds", "86400", "int", "JWT token expiration (24 hours)"},
		{"consent.token.expiration.seconds", "3600", "int", "Consent token expiration (1 hour)"},
		{"auth.context.retention.days", "30", "int", "Authentication context retention period"},
		{"rate.limiting.enabled", "true", "bool", "Enable rate limiting"},
		{"rate.limiting.anonymous.per.minute", "100", "int", "Anonymous requests per minute"},
		{"rate.limiting.authenticated.per.minute", "1000", "int", "Authenticated requests per minute"},
	}

	for _, config := range defaults {
		if err := cs.SetConfig(config.key, config.value, config.configType, config.description); err != nil {
			return err
		}
	}

	return nil
}
