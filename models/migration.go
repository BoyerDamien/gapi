package models

import (
	"log"

	"gorm.io/gorm"
)

func MigrateAll(db *gorm.DB) error {
	log.Println("Models migration in progress...")
	return db.AutoMigrate(
		&Media{},
		&User{},
		&Portfolio{},
		&Offer{},
		&Tag{},
	)
}
