package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() {

	// postgres://:@/
	dsn := "host=mouse.db.elephantsql.com user=ixwniyxv password=JRBipBbQtmMw36ZNr703S-PcV-o9Rqyq dbname=ixwniyxv port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
