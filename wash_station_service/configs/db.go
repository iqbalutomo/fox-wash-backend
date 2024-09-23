package config

import (
	"fmt"
	"os"
	"wash_station_service/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	dsn := os.Getenv("POSTGRES_URI")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %v", err)
	}

	if err := db.AutoMigrate(&models.Wash{}, &models.Detailing{}); err != nil {
		return nil, fmt.Errorf("failed to migrating database: %v", err)
	}

	return db, nil
}
