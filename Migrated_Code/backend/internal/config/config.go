package config

import (
	"os"
)

type Config struct {
	Port     string
	Database DatabaseConfig
	API      APIConfig
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
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
