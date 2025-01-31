package database

import (
	"HuCap/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(cfg *config.Config) (*gorm.DB, error) {

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,cfg.DBSSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn),&gorm.Config{})

	if err != nil {
		return nil , err
	}

	log.Println("Database connection successfully.")
	RunMigrations(db)
	return db, nil
}