package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	AppName string `env:"APP_NAME" env-required:"true"`
	AppPort string `env:"APP_PORT" env-default:"8080"`

	DBHost     string `env:"DB_HOST" env-required:"true"`
	DBPort     string `env:"DB_PORT" env-default:"5432"`
	DBUser     string `env:"DB_USER" env-required:"true"`
	DBPassword string `env:"DB_PASS" env-required:"true"`
	DBName     string `env:"DB_NAME" env-required:"true"`
	DBSSLMode  string `env:"DB_SSLMODE" env-default:"disable"`

	JWTSecret      string `env:"JWT_SECRET" env-required:"true"`
	JWTExpireHours int    `env:"JWT_EXPIRE_HOURS" env-default:"24"`
}

func Load() *Config {
	_ = godotenv.Load(".ENV")

	var cfg Config

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatalf("cannot read env: %v", err)
	}

	return &cfg
}
