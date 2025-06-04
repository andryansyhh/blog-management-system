package cmd

import (
	"blog-management-system/internal/domain/model"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewClientDatabase(cfg *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	// migration init
	InitMigration(db)

	return db, nil
}

func InitMigration(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.User{},
		&model.Content{},
		&model.Category{},
	)
}
