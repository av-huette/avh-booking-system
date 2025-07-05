package logger

import (
	"fmt"
	"github.com/av-huette/avh-booking-system/config"
	"github.com/lmittmann/tint"
	"log/slog"
	"os"
	"time"
)

func CreateLogger() *slog.Logger {
	options := &tint.Options{
		Level:      config.LoadConfig().LogLevel,
		TimeFormat: time.DateTime,
	}

	log := slog.New(tint.NewHandler(os.Stdout, options))
	log.Debug(fmt.Sprintf("Log level: %s", options.Level))
	
	return log
}
