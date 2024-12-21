package main

import (
	"os"

	clog "github.com/charmbracelet/log"
	"github.com/cptpie/tsloo/logging"
)

func main() {
	// config := backend.Options{
	// 	Debug:  debug,
	// 	Listen: ":8090",
	// }
	//

	logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		clog.Error("Could not create/open logfile: %v", err.Error())
		os.Exit(1)
	}

	log := logging.New(*logFile)
	log.Warn("hello world!")

	// var db string
	// db, err := database.GetDatabase()
	if err != nil {
		log.Error("Could not open database connection: %v", err.Error())
		os.Exit(1)
	}

	// var backend string
	// backend, err := backend.New(db, log)
}
