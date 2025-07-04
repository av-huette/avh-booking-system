package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: " + err.Error())
	}

	// set up file server
	staticPath := os.Getenv("AVHBS_FRONTEND_PATH")
	fileServer := http.FileServer(http.Dir(staticPath))

	// set up routes
	mux := http.NewServeMux()
	mux.Handle("GET /", fileServer)

	// set up web server
	server := &http.Server{
		Addr:           ":8081",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Print("Starting server on :8081")
	err = server.ListenAndServe()
	log.Fatal(err)
}
