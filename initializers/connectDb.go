package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
