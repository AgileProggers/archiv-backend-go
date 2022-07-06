package database

import (
	"github.com/AgileProggers/archiv-backend-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() {
	dsn := "host=localhost user=test password=test dbname=test port=5432 sslmode=disable TimeZone=Europe/Berlin"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(&models.Vod{}, &models.Game{}, &models.Creator{}, &models.Clip{})
	DB = database
}
