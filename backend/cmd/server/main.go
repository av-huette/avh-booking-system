package main

import (
	"github.com/av-huette/avh-booking-system/config"
	"github.com/av-huette/avh-booking-system/internal/logger"
	"log/slog"
	"os"
)

type application struct {
	conf *config.Config
	log  *slog.Logger
}

func main() {
	app := &application{
		conf: config.LoadConfig(),
		log:  logger.CreateLogger(),
	}

	err := app.serveHTTP()
	if err != nil {
		app.log.Error(err.Error())
		os.Exit(1)
	}
}
