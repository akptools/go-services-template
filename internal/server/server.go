package server

import (
	"database/sql"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
)

// Defines the Server struct contaning a listenAddr and a pointer to db
type Server struct {
	listenAddr string
	db         *sql.DB
}

// Initializes a new server object and returns it
func NewServer(listenAddr string, db *sql.DB) *Server {
	return &Server{
		listenAddr: listenAddr,
		db:         db,
	}
}

// Registers the routes and starts the server
func (s *Server) Run() {
	r := chi.NewRouter()
	registerMiddleware(r)
	registerRoutes(r)

	slog.Info("Listening on", "port", strings.TrimPrefix(s.listenAddr, ":"))
	err := http.ListenAndServe(s.listenAddr, r)
	if err != nil {
		slog.Error("An error occured while starting the server", "err:", err.Error())
		os.Exit(1)
	}
}
