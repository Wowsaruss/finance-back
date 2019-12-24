package config

import (
	"os"
	"strconv"
)

// Config ...
type Config struct {
	HostPort       string
	DBHost         string
	DBPort         int
	DBUser         string
	DBPassword     string
	DBName         string
	MongoDBURI     string
	PlaidClientID  string
	PlaidSecret    string
	PlaidPublicKey string
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		HostPort:       getEnv("PORT", ":10000"),
		DBHost:         getEnv("DB_HOST", ""),
		DBUser:         getEnv("DB_USER", ""),
		DBPassword:     getEnv("DB_PASSWORD", ""),
		DBName:         getEnv("DB_NAME", ""),
		DBPort:         getEnvAsInt("DB_PORT", 5432),
		MongoDBURI:     getEnv("MONGODB_URI", ""),
		PlaidClientID:  getEnv("PLAID_CLIENT_ID", ""),
		PlaidSecret:    getEnv("PLAID_SECRET", ""),
		PlaidPublicKey: getEnv("PLAID_PUBLIC_KEY", ""),
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
