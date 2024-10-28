package server

import (
	"github.com/akptools/go-services-template/internal/handler"
	"github.com/akptools/go-services-template/internal/util"
	"github.com/go-chi/chi/v5"
)

func registerRoutes(r *chi.Mux) {
	r.Get("/healthcheck", util.Make(handler.Healthcheck))
}
