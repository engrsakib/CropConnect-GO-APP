package db

import (
	"log"

	"gorm.io/gorm"
	"github.com/engrsakib/erp-system/internal/models"
)


func Migrate(db *gorm.DB) error {
	log.Println("🔄 Running Database Migration...")

	if err := db.Exec("CREATE EXTENSION IF NOT EXISTS postgis;").Error; err != nil {
		log.Printf("❌ Failed to enable PostGIS extension: %v", err)
		return err
	}
	
	err := db.AutoMigrate(
		&models.User{},           
		&models.UserPermission{}, 
		&models.Faq{},       
        &models.FaqAnswer{},
		&models.GroupType{},
		&models.Group{},
		&models.Member{},
		&models.Geofence{},
		&models.GeofencePoint{},
		&models.Crop{},
	)

	if err != nil {
		log.Printf("⚠️ Warning: AutoMigrate failed: %v", err)
		return err
	}

	log.Println("✅ Database Migration Successful")
	return nil
}