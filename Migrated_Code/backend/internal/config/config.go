package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	Port           string
	Database       DatabaseConfig
	JWT            JWTConfig
	Authentication AuthenticationConfig
	Security       SecurityConfig
	RateLimit      RateLimitConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type JWTConfig struct {
	Secret     string
	Expiration time.Duration
}

type AuthenticationConfig struct {
	DirectLoginEnabled        bool
	DirectLoginTokenExpiry    time.Duration
	OAuthEnabled              bool
	OAuthTokenExpiry          time.Duration
	OAuth2Enabled             bool
	OAuth2TokenExpiry         time.Duration
	GatewayLoginEnabled       bool
	MaxBadLoginAttempts       int
	UserLockDuration          time.Duration
	BcryptCost                int
	ContextRetentionDays      int
	ConsentTokenExpiry        time.Duration
}

type SecurityConfig struct {
	RequireClientCertificate  bool
	CertificateKeystorePath   string
	CertificateKeystorePass   string
	PasswordMinLength         int
	PasswordRequireSpecial    bool
	SessionTimeout            time.Duration
}

type RateLimitConfig struct {
	Enabled                   bool
	AnonymousPerMinute        int
	AnonymousPerHour          int
	AuthenticatedPerMinute    int
	AuthenticatedPerHour      int
	ConsumerSpecificEnabled   bool
}

func Load() *Config {
	return &Config{
		Port: getEnv("PORT", "8080"),
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", ""),
			DBName:   getEnv("DB_NAME", "obp_api"),
			SSLMode:  getEnv("DB_SSL_MODE", "disable"),
		},
		JWT: JWTConfig{
			Secret:     getEnv("JWT_SECRET", "your-secret-key-change-in-production"),
			Expiration: getDurationEnv("JWT_EXPIRATION", 24*time.Hour),
		},
		Authentication: AuthenticationConfig{
			DirectLoginEnabled:        getBoolEnv("DIRECT_LOGIN_ENABLED", true),
			DirectLoginTokenExpiry:    getDurationEnv("DIRECT_LOGIN_TOKEN_EXPIRY", 4*7*24*time.Hour),
			OAuthEnabled:              getBoolEnv("OAUTH_ENABLED", true),
			OAuthTokenExpiry:          getDurationEnv("OAUTH_TOKEN_EXPIRY", 1*time.Hour),
			OAuth2Enabled:             getBoolEnv("OAUTH2_ENABLED", true),
			OAuth2TokenExpiry:         getDurationEnv("OAUTH2_TOKEN_EXPIRY", 1*time.Hour),
			GatewayLoginEnabled:       getBoolEnv("GATEWAY_LOGIN_ENABLED", false),
			MaxBadLoginAttempts:       getIntEnv("MAX_BAD_LOGIN_ATTEMPTS", 5),
			UserLockDuration:          getDurationEnv("USER_LOCK_DURATION", 30*time.Minute),
			BcryptCost:                getIntEnv("BCRYPT_COST", 12),
			ContextRetentionDays:      getIntEnv("CONTEXT_RETENTION_DAYS", 30),
			ConsentTokenExpiry:        getDurationEnv("CONSENT_TOKEN_EXPIRY", 1*time.Hour),
		},
		Security: SecurityConfig{
			RequireClientCertificate: getBoolEnv("REQUIRE_CLIENT_CERTIFICATE", false),
			CertificateKeystorePath:  getEnv("CERTIFICATE_KEYSTORE_PATH", ""),
			CertificateKeystorePass:  getEnv("CERTIFICATE_KEYSTORE_PASSWORD", ""),
			PasswordMinLength:        getIntEnv("PASSWORD_MIN_LENGTH", 8),
			PasswordRequireSpecial:   getBoolEnv("PASSWORD_REQUIRE_SPECIAL", true),
			SessionTimeout:           getDurationEnv("SESSION_TIMEOUT", 30*time.Minute),
		},
		RateLimit: RateLimitConfig{
			Enabled:                 getBoolEnv("RATE_LIMIT_ENABLED", true),
			AnonymousPerMinute:      getIntEnv("RATE_LIMIT_ANONYMOUS_PER_MINUTE", 100),
			AnonymousPerHour:        getIntEnv("RATE_LIMIT_ANONYMOUS_PER_HOUR", 1000),
			AuthenticatedPerMinute:  getIntEnv("RATE_LIMIT_AUTHENTICATED_PER_MINUTE", 1000),
			AuthenticatedPerHour:    getIntEnv("RATE_LIMIT_AUTHENTICATED_PER_HOUR", 10000),
			ConsumerSpecificEnabled: getBoolEnv("RATE_LIMIT_CONSUMER_SPECIFIC", true),
		},
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func getIntEnv(key string, fallback int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return fallback
}

func getBoolEnv(key string, fallback bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return fallback
}

func getDurationEnv(key string, fallback time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}
	return fallback
}
