package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PostgresDSN string
}

var AppConfig *Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	AppConfig = &Config{
		PostgresDSN: os.Getenv("POSTGRES_DSN"),
	}
}

func GetPostgresDSN() string {
	return AppConfig.PostgresDSN
}
