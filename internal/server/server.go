package server

import (
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/akptools/go-services-template/internal/db"
	"github.com/go-chi/chi/v5"
)

// Defines the Server struct contaning a listenAddr and a pointer to db
type Server struct {
	listenAddr string
	db         db.Database
}

// Initializes a new server object and returns it
func NewServer(listenAddr string, db db.Database) *Server {
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
