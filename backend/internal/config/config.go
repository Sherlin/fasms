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

func GetDSN() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
    user := os.Getenv("DB_USER")
    pass := os.Getenv("DB_PASS")
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    dbname := os.Getenv("DB_NAME")
    
    if user == "" || pass == "" || host == "" || port == "" || dbname == "" {
	
        log.Fatal("Missing database credentials in environment variables.")
    }

    return user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
}