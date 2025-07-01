package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	staticPath := os.Getenv("AVHBS_FRONTEND_PATH")
	if staticPath == "" {
		panic("Environment variable AVHBS_FRONTEND_PATH is not set. Make sure to set it to `<path to avh-booking-system>/client/dist`")
	}

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir(staticPath))
	mux.Handle("GET /{$}", fileServer)

	server := &http.Server{
		Addr:           ":8081",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Print("Starting server on :8081")
	err := server.ListenAndServe()
	log.Fatal(err)
}
