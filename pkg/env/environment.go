package env

import (
	"github.com/AgileProggers/archiv-backend-go/pkg/logger"
	"github.com/joho/godotenv"
	"os"
)

var (
	PostgresDatabase string
)

func init() {
	// Load will not override existing variables
	// However godotenv.Overload does
	err := godotenv.Load()
	if err != nil {
		logger.Error.Fatalln("Unable to load environment:", err)
	}

	PostgresDatabase = os.Getenv("POSTGRES_DATABASE")

}
