package database

import (
	"fmt"
	"log"

	"github.com/Yogi-1996/notes-backend/internal/config"
	"github.com/Yogi-1996/notes-backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.DB.Host, cfg.DB.User, cfg.DB.Password, cfg.DB.Name, cfg.DB.Port, cfg.DB.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to DB: %w", err)
	}

	if err := db.AutoMigrate(
		&models.User{},
		&models.Note{},
	); err != nil {
		log.Fatal("Failed to automigrate database:", err)
	}

	log.Println("Database connected and migrated successfully")
	return db, nil
}
