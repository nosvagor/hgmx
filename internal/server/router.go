package server

import (
	"net/http"
)

// NewRouter creates a new ServeMux and registers the application handlers.
func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// Register handlers from handlers.go
	mux.HandleFunc("/", Index)

	// TODO: Add routes for serving static assets (CSS, JS) later

	return mux
}
