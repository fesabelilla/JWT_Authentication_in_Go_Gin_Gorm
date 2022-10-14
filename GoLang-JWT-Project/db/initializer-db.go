package db

import "gin-mongo-api/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
