package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Port            string
	Environment     string
	AllowedOrigins  string
	RateLimit       int
	RateLimitWindow time.Duration
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	return &Config{
		Port:            getEnv("PORT", "3000"),
		Environment:     getEnv("ENVIRONMENT", "development"),
		AllowedOrigins:  getEnv("ALLOWED_ORIGINS", "*"),
		RateLimit:       getEnvAsInt("RATE_LIMIT", 100),
		RateLimitWindow: getEnvAsDuration("RATE_LIMIT_WINDOW", 1*time.Minute),
		ReadTimeout:     getEnvAsDuration("READ_TIMEOUT", 10*time.Second),
		WriteTimeout:    getEnvAsDuration("WRITE_TIMEOUT", 10*time.Second),
	}
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Printf("Invalid value for %s, using default: %d", key, defaultValue)
		return defaultValue
	}

	return value
}

func getEnvAsDuration(key string, defaultValue time.Duration) time.Duration {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}

	value, err := time.ParseDuration(valueStr)
	if err != nil {
		log.Printf("Invalid duration for %s, using default: %v", key, defaultValue)
		return defaultValue
	}

	return value
}
