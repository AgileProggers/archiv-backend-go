package database

import (
	"fmt"
	"github.com/AgileProggers/archiv-backend-go/pkg/env"
	"github.com/AgileProggers/archiv-backend-go/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"time"
)

var database *gorm.DB

func init() {
	// Connect DB

	db, err := gorm.Open(postgres.Open(env.PostgresDatabase), &gorm.Config{
		Logger: gormLogger.New(logger.Info, gormLogger.Config{}),
	})
	if err != nil {
		logger.Error.Fatalln("Unable to connect to database:", err)
	}
	err = db.AutoMigrate(&Vod{}, &Game{}, &Creator{}, &Clip{})
	if err != nil {
		logger.Error.Fatalln("Unable to auto migrate database:", err)
	}
	database = db
}

func Ping() (time.Duration, error) {
	start := time.Now()
	db, err := database.DB()
	if err != nil {
		return 0, fmt.Errorf("gorm.DB get database: %v", err)
	}
	err = db.Ping()
	if err != nil {
		return 0, fmt.Errorf("databse ping: %v", err)
	}
	return time.Now().Sub(start), nil
}

func Close() error {
	db, err := database.DB()
	if err != nil {
		return fmt.Errorf("gorm.DB get database: %v", err)
	}
	return db.Close()
}
func Raw(sql string, values ...interface{}) (tx *gorm.DB) {
	return database.Raw(sql, values...)
}
