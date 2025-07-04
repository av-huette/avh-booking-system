package server

import (
	"net/http"
	"os"
)

func SetUpMux() http.Handler {
	// set up file server
	staticPath := os.Getenv("AVHBS_FRONTEND_PATH")
	fileServer := http.FileServer(http.Dir(staticPath))

	// set up routes
	mux := http.NewServeMux()
	mux.Handle("GET /", fileServer)

	return mux
}
