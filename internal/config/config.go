package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, reading environment variables directly")
	}

	//fjhfnhfjhfjhgf

	cfg := &Config{
		DBHost: os.Getenv("DB_HOST"),
		DBPort: os.Getenv("DB_PORT"),
		DBUser: os.Getenv("DB_USER"),
		DBPass: os.Getenv("DB_PASS"),
		DBName: os.Getenv("DB_NAME"),
	}

	// sanity check
	if cfg.DBHost == "" || cfg.DBPort == "" || cfg.DBUser == "" || cfg.DBPass == "" || cfg.DBName == "" {
		log.Fatal("Some DB environment variables are missing")
	}

	return cfg
}
