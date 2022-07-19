package database

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/AgileProggers/archiv-backend-go/pkg/ent"
	"github.com/AgileProggers/archiv-backend-go/pkg/env"
	"github.com/AgileProggers/archiv-backend-go/pkg/logger"
	_ "github.com/lib/pq"

	"time"
)

var (
	client *ent.Client
	driver *sql.Driver
)

func init() {
	logger.Info.Println("Initializing Postgres database connection")
	var err error
	driver, err = sql.Open("postgres", env.PostgresDatabase)
	if err != nil {
		logger.Error.Fatalln("Failed opening connection to Postgres:", err)
	}
	// Get the underlying sql.DB object of the driver.
	db := driver.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	client = ent.NewClient(ent.Driver(driver)) //.Debug()
	err = autoMigrate()
	if err != nil {
		logger.Error.Fatalln("Failed to auto migrate database:", err)
	}
	logger.Info.Println("Migrated Postgres database")
}

func autoMigrate() error {
	return client.Schema.Create(context.Background())
}

func rollback(errToReturn error, tx *ent.Tx) error {
	err := tx.Rollback()
	if err != nil {
		return err
	}
	return errToReturn
}

func Ping() (time.Duration, error) {
	start := time.Now()
	err := driver.DB().Ping()
	return time.Now().Sub(start), err
}

func Close() error {
	return client.Close()
}
