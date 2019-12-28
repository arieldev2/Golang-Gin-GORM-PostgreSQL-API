package migrations

import (
	"gin_test/models"

	"gin_test/database"
)

func InitialMigration() {
	db := database.Connection()

	defer db.Close()

	db.AutoMigrate(&models.Post{})
}
