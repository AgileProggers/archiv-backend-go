package signals

import (
	"github.com/AgileProggers/archiv-backend-go/pkg/database"
	"github.com/AgileProggers/archiv-backend-go/pkg/logger"
	"github.com/AgileProggers/archiv-backend-go/pkg/router"
	"os"
	"os/signal"
	"syscall"
)

var done = make(chan bool, 1)

func ListenForSigterm() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		logger.Debug.Println("Received signal", sig)
		err := router.Shutdown()
		if err != nil {
			logger.Error.Println("Could not stop router:", err)
		}
		logger.Info.Println("Closed router")

		err = database.Close()
		if err != nil {
			logger.Error.Println("Could not close database connection:", err)
		}
		logger.Info.Println("Closed database")

		logger.Info.Println("Stopping agile archive backend")
		done <- true
	}()

}

func WaitForCleanup() {
	<-done
}
