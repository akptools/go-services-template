package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func registerMiddleware(r *chi.Mux) {
	// Generates a random integer and injects into the request
	r.Use(middleware.RequestID)
	// Logs the request and response with some useful data
	r.Use(middleware.Logger)
	// Recovers the application from panics and returns a 500 internal server error
	r.Use(middleware.Recoverer)
}
