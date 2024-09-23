package migrates

import (
	"log"
	"user_service/models"
	"user_service/utils"

	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB) {
	roles := []models.Role{
		{
			ID:   utils.UserRoleID,
			Name: utils.UserRole,
		},
		{
			ID:   utils.WasherRoleID,
			Name: utils.WasherRole,
		},
		{
			ID:   utils.AdminRoleID,
			Name: utils.AdminRole,
		},
	}

	for _, role := range roles {
		var existingRole models.Role

		result := db.First(&existingRole, "name = ?", role.Name)
		if result.Error != nil && result.Error == gorm.ErrRecordNotFound {
			if err := db.Create(&role).Error; err != nil {
				log.Printf("failed to created role %s: %v", role.Name, err)
			} else {
				log.Printf("role %s created", role.Name)
			}
		} else {
			log.Printf("role %s already exists", role.Name)
		}
	}

	log.Println("seeding roles completed.")
}

func SeedWasherStatuses(db *gorm.DB) {
	statuses := []models.WasherStatus{
		{
			ID:     utils.AvailableWasherStatusID,
			Status: utils.AvailableStatus,
		},
		{
			ID:     utils.WashingWasherStatusID,
			Status: utils.WashingStatus,
		},
		{
			ID:     utils.InActiveWasherStatusID,
			Status: utils.InActiveStatus,
		},
	}

	for _, washer := range statuses {
		var existingStatus models.WasherStatus

		result := db.First(&existingStatus, "status = ?", washer.Status)
		if result.Error != nil && result.Error == gorm.ErrRecordNotFound {
			if err := db.Create(&washer).Error; err != nil {
				log.Printf("failed to create washer status %s: %v", washer.Status, err)
			} else {
				log.Printf("washer status %s created", washer.Status)
			}
		} else {
			log.Printf("washer status %s already exists", washer.Status)
		}
	}

	log.Println("seeding washer_statuses completed.")
}
