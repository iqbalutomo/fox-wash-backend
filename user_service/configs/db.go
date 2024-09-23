package config

import (
	"fmt"
	"os"
	"user_service/migrates"
	"user_service/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	dsn := os.Getenv("POSTGRES_URI")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %v", err)
	}

	if err := db.AutoMigrate(&models.Role{}, &models.User{}, &models.EmailVerification{}, &models.WasherStatus{}, &models.Washer{}); err != nil {
		return nil, fmt.Errorf("failed to migrating database: %v", err)
	}

	migrates.SeedRoles(db)
	migrates.SeedWasherStatuses(db)

	return db, nil
}
