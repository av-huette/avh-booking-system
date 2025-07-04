package main

import (
	"github.com/av-huette/avh-booking-system/internal/server"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"time"
)

func main() {
	// load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: " + err.Error())
	}

	// set up web server
	webServer := &http.Server{
		Addr:           ":8081",
		Handler:        server.SetUpMux(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Print("Starting server on :8081")
	err = webServer.ListenAndServe()
	log.Fatal(err)
}
