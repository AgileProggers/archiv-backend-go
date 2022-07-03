package database

import (
	"github.com/AgileProggers/archiv-backend-go/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(&models.Vod{})

	DB = database
}
