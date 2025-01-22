package config

import (
	"os"
	log "github.com/sirupsen/logrus"
	"github.com/joho/godotenv"
)

// Config holds application configuration.
type Config struct {
	Port string
}

// LoadConfig loads configuration from environment variables.
func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080" // Default port
	}

	return &Config{
		Port: port,
	}, nil
}
