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
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.DBSSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to DB: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql.DB: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("database ping failed: %w", err)
	}

	if err := db.AutoMigrate(
		&models.User{},
		&models.Note{},
	); err != nil {
		return nil, fmt.Errorf("failed to automigrate database: %w", err)
	}

	log.Println("Database connected and migrated successfully")
	return db, nil
}
