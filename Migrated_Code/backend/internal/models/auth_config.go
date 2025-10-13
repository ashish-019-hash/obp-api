package models

import (
	"time"
)

type AuthenticationConfig struct {
	ID          int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	ConfigKey   string    `json:"config_key" gorm:"uniqueIndex;size:255;not null"`
	ConfigValue string    `json:"config_value" gorm:"type:text;not null"`
	ConfigType  string    `json:"config_type" gorm:"size:50;not null"`
	Description string    `json:"description" gorm:"type:text"`
	IsActive    bool      `json:"is_active" gorm:"default:true"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ConsumerRateLimit struct {
	ID                int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	ConsumerID        string    `json:"consumer_id" gorm:"index;size:255;not null"`
	RequestsPerMinute int       `json:"requests_per_minute" gorm:"default:100"`
	RequestsPerHour   int       `json:"requests_per_hour" gorm:"default:1000"`
	RequestsPerDay    int       `json:"requests_per_day" gorm:"default:10000"`
	IsActive          bool      `json:"is_active" gorm:"default:true"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	
	Consumer          *Consumer `json:"consumer,omitempty" gorm:"foreignKey:ConsumerID;references:ConsumerID"`
}

type SecuritySettings struct {
	ID           int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	SettingKey   string    `json:"setting_key" gorm:"uniqueIndex;size:255;not null"`
	SettingValue string    `json:"setting_value" gorm:"type:text;not null"`
	SettingType  string    `json:"setting_type" gorm:"size:50;not null"`
	Description  string    `json:"description" gorm:"type:text"`
	IsActive     bool      `json:"is_active" gorm:"default:true"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type TokenConfiguration struct {
	ID                    int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	TokenType             string    `json:"token_type" gorm:"index;size:50;not null"`
	ExpirationSeconds     int64     `json:"expiration_seconds" gorm:"not null"`
	RefreshEnabled        bool      `json:"refresh_enabled" gorm:"default:false"`
	RefreshExpirationSecs int64     `json:"refresh_expiration_seconds" gorm:"default:0"`
	IsActive              bool      `json:"is_active" gorm:"default:true"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}

func (AuthenticationConfig) TableName() string {
	return "authentication_configs"
}

func (ConsumerRateLimit) TableName() string {
	return "consumer_rate_limits"
}

func (SecuritySettings) TableName() string {
	return "security_settings"
}

func (TokenConfiguration) TableName() string {
	return "token_configurations"
}

func NewAuthenticationConfig(key, value, configType, description string) *AuthenticationConfig {
	return &AuthenticationConfig{
		ConfigKey:   key,
		ConfigValue: value,
		ConfigType:  configType,
		Description: description,
		IsActive:    true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func NewConsumerRateLimit(consumerID string, perMinute, perHour, perDay int) *ConsumerRateLimit {
	return &ConsumerRateLimit{
		ConsumerID:        consumerID,
		RequestsPerMinute: perMinute,
		RequestsPerHour:   perHour,
		RequestsPerDay:    perDay,
		IsActive:          true,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}
}

func NewSecuritySettings(key, value, settingType, description string) *SecuritySettings {
	return &SecuritySettings{
		SettingKey:   key,
		SettingValue: value,
		SettingType:  settingType,
		Description:  description,
		IsActive:     true,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

func NewTokenConfiguration(tokenType string, expirationSeconds int64, refreshEnabled bool, refreshExpirationSecs int64) *TokenConfiguration {
	return &TokenConfiguration{
		TokenType:             tokenType,
		ExpirationSeconds:     expirationSeconds,
		RefreshEnabled:        refreshEnabled,
		RefreshExpirationSecs: refreshExpirationSecs,
		IsActive:              true,
		CreatedAt:             time.Now(),
		UpdatedAt:             time.Now(),
	}
}
