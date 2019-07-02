package config

import (
	"os"
)

// Config ...
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		Host:     getEnv("DB_HOST", ""),
		Port:     getEnv("DB_PORT", ""),
		User:     getEnv("DB_USER", ""),
		Password: getEnv("DB_PASSWORD", ""),
		DBName:   getEnv("DB_NAME", ""),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
