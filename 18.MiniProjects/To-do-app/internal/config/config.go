package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	Port        string
}

func MustLoad() (*Config, error) {
	var err error = godotenv.Load()
	if err != nil {
		log.Fatal("failed to load env: %w", err)
	}

	var cfg *Config = &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Port:        os.Getenv("PORT"),
	}

	return cfg, nil
}
