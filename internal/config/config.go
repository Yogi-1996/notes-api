package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name string
		Port string
		Env  string
	}
	DB struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
		SSLMode  string
	}
	JWT struct {
		Secret      string
		ExpireHours int
	}
}

func Load() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Unable to decode config: %v", err)
	}

	if cfg.DB.Host == "" || cfg.DB.User == "" {
		log.Fatal("Database configuration missing")
	}

	if cfg.App.Port == "" {
		cfg.App.Port = "8080"
	}

	return &cfg
}
