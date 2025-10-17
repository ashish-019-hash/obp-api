package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port     string
	Database DatabaseConfig
	API      APIConfig
	Auth     AuthConfig
}

type DatabaseConfig struct {
	Type string
	Host string
	Port string
	Name string
	User string
	Pass string
}

type APIConfig struct {
	Version string
	Title   string
	BaseURL string
}

type AuthConfig struct {
	AllowDirectLogin           bool
	DirectLoginSecret          string
	DirectLoginTokenExpiration int
	AllowOAuth1                bool
	OAuth1TokenExpiration      int
	AllowOAuth2                bool
	OAuth2JWKSetURL            string
	AllowGatewayLogin          bool
	GatewayTokenSecret         string
	AllowDAuth                 bool
	MaxBadLoginAttempts        int
	UserLockDurationSeconds    int
}

func Load() *Config {
	return &Config{
		Port: getEnv("PORT", "8080"),
		Database: DatabaseConfig{
			Type: getEnv("DB_TYPE", "memory"),
			Host: getEnv("DB_HOST", "localhost"),
			Port: getEnv("DB_PORT", "5432"),
			Name: getEnv("DB_NAME", "obp_api"),
			User: getEnv("DB_USER", "postgres"),
			Pass: getEnv("DB_PASS", ""),
		},
		API: APIConfig{
			Version: getEnv("API_VERSION", "5.1.0"),
			Title:   getEnv("API_TITLE", "Open Bank Project API"),
			BaseURL: getEnv("API_BASE_URL", "http://localhost:8080"),
		},
		Auth: AuthConfig{
			AllowDirectLogin:           getBoolEnv("ALLOW_DIRECT_LOGIN", true),
			DirectLoginSecret:          getEnv("DIRECT_LOGIN_SECRET", "change-this-secret-in-production-use-strong-random-value"),
			DirectLoginTokenExpiration: getIntEnv("DIRECT_LOGIN_TOKEN_EXPIRATION", 2419200),
			AllowOAuth1:                getBoolEnv("ALLOW_OAUTH1_LOGIN", false),
			OAuth1TokenExpiration:      getIntEnv("OAUTH1_TOKEN_EXPIRATION", 3600),
			AllowOAuth2:                getBoolEnv("ALLOW_OAUTH2_LOGIN", false),
			OAuth2JWKSetURL:            getEnv("OAUTH2_JWK_SET_URL", ""),
			AllowGatewayLogin:          getBoolEnv("ALLOW_GATEWAY_LOGIN", false),
			GatewayTokenSecret:         getEnv("GATEWAY_TOKEN_SECRET", ""),
			AllowDAuth:                 getBoolEnv("ALLOW_DAUTH", false),
			MaxBadLoginAttempts:        getIntEnv("MAX_BAD_LOGIN_ATTEMPTS", 5),
			UserLockDurationSeconds:    getIntEnv("USER_LOCK_DURATION_SECONDS", 1800),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getBoolEnv(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		boolValue, err := strconv.ParseBool(value)
		if err == nil {
			return boolValue
		}
	}
	return defaultValue
}

func getIntEnv(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		intValue, err := strconv.Atoi(value)
		if err == nil {
			return intValue
		}
	}
	return defaultValue
}
