package database

import (
	"github.com/c9b3rd3vi1/NetGuardIQ/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB is the global database connection
var DB *gorm.DB

// ConnectDB initializes the database connection
func ConnectDB() error {
	var err error
	DB, err = gorm.Open(sqlite.Open("netguardiq.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database: " + err.Error())
		//return err
	}

	// Migrate the schema
	err = DB.AutoMigrate(&models.Target{}, &models.Campaign{})
	if err != nil {
		return err
	}

	return nil
}