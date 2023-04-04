package initializers

import "userauth/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
