package main

import (
	"os"
	"strconv"
)

// Config ...
type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		Host:     getEnv("DB_HOST", ""),
		User:     getEnv("DB_USER", ""),
		Password: getEnv("DB_PASSWORD", ""),
		DBName:   getEnv("DB_NAME", ""),
		Port:     getEnvAsInt("DB_PORT", 5432),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}
