package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port    string
	Service string
	Version string
}

func Load() *Config {
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	return &Config{
		Port:    port,
		Service: "KageVault API",
		Version: "1.0.0",
	}
}
