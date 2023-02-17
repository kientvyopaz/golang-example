package migrations

import (
	"golang-example/config"
	"golang-example/models"
)

func AutoMigrate() {
	db := config.DB

	// Run migration for User model
	db.AutoMigrate(&models.User{})
}
