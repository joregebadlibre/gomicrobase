package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	POSTGRES_HOST     string
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_DB       string
}

var AppConfig *Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	AppConfig = &Config{
		POSTGRES_HOST:     os.Getenv("POSTGRES_HOST"),
		POSTGRES_USER:     os.Getenv("POSTGRES_USER"),
		POSTGRES_PASSWORD: os.Getenv("POSTGRES_PASSWORD"),
		POSTGRES_DB:       os.Getenv("POSTGRES_DB"),
	}
}

func GetPostgresDSN() string {
	log.Println("Variables de Entorno: ", AppConfig.POSTGRES_HOST, AppConfig.POSTGRES_USER, AppConfig.POSTGRES_PASSWORD, AppConfig.POSTGRES_DB)
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", AppConfig.POSTGRES_HOST, AppConfig.POSTGRES_USER, AppConfig.POSTGRES_PASSWORD, AppConfig.POSTGRES_DB)
}
