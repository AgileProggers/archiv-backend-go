package main

import (
	"net/http"
	"os"

	"github.com/AgileProggers/archiv-backend-go/pkg/database"
	"github.com/AgileProggers/archiv-backend-go/pkg/logger"
	"github.com/AgileProggers/archiv-backend-go/pkg/router"
	"github.com/AgileProggers/archiv-backend-go/pkg/signals"
)

func main() {
	signals.ListenForSigterm()

	time, err := database.Ping()
	if err != nil {
		logger.Error.Fatalln("Could not ping the database:", err)
	}
	logger.Info.Println("Postgres database delay is", time)

	logger.Info.Println("Starting to listen on port 8080 on PID", os.Getpid())
	err = router.Listen()
	if err != nil && err != http.ErrServerClosed {
		logger.Error.Fatalln("Could not start the router:", err)
	}

	signals.WaitForCleanup()
}
