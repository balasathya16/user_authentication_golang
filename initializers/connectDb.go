package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	var err error
	// postgres://:@/
	dsn := "host=mouse.db.elephantsql.com user=ixwniyxv password=JRBipBbQtmMw36ZNr703S-PcV-o9Rqyq dbname=ixwniyxv port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to db")
	}
}
