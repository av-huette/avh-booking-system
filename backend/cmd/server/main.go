package main

import (
	"fmt"
	"github.com/av-huette/avh-booking-system/config"
	"github.com/av-huette/avh-booking-system/internal/server"
	"log"
	"net/http"
	"time"
)

func main() {
	// load config
	conf := config.LoadConfig()

	// set up web server
	webServer := &http.Server{
		Addr:           fmt.Sprintf(":%d", conf.HttpPort),
		Handler:        server.SetUpMux(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Print("Starting server on :8081")
	err := webServer.ListenAndServe()
	log.Fatal(err)
}
