package database

import (
	"HuCap/base/models"
	"log"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {

	err := db.AutoMigrate(
		&models.User{},
	)
	
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migration completed.")
}