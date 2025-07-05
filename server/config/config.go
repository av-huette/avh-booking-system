package config

import (
	"github.com/joho/godotenv"
	"log"
	"log/slog"
	"os"
	"strconv"
)

type Config struct {
	LogLevel     slog.Level
	HttpPort     int
	FrontendPath string
}

func LoadConfig() *Config {
	// load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: " + err.Error())
	}

	conf := &Config{}
	conf.LogLevel = getLogLevel("AVHBS_LOG_LEVEL", slog.LevelInfo)
	conf.HttpPort = getInt("AVHBS_HTTP_PORT", 8081)
	conf.FrontendPath = getString("AVHBS_FRONTEND_PATH", "")

	return conf
}

func getString(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	return value
}

func getInt(key string, defaultValue int) int {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}

	return intValue
}

func getBool(key string, defaultValue bool) bool {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		panic(err)
	}

	return boolValue
}

func getLogLevel(key string, defaultValue slog.Level) slog.Level {
	value, exists := os.LookupEnv(key)
	if !exists {
		return slog.LevelInfo
	}

	m := make(map[string]slog.Level)
	m["DEBUG"] = slog.LevelDebug
	m["INFO"] = slog.LevelInfo
	m["WARN"] = slog.LevelWarn
	m["ERROR"] = slog.LevelError

	logLevel, exists := m[value]
	if exists {
		return logLevel
	}
	return defaultValue
}
