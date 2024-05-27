package server

import (
	"log/slog"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"chi-learn/internals/database"
)

type Service struct {
	logger *slog.Logger
	Router *chi.Mux
	DB     database.Service
}

func New(logger *slog.Logger, db database.Service) *Service {
	s := &Service{
		logger: logger,
		Router: chi.NewRouter(),
		DB:     db,
	}

	// INFO: Define all middlewares before mounting the handlers
	s.Router.Use(middleware.RequestID)
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)

	s.Router.Use(middleware.CleanPath)

	// INFO: Mount all the handlers
	s.MountHandlers()

	return s
}

func (s *Service) MountHandlers() {
	s.Router.Get("/", make(s.Home))

	s.Router.Post("/companies", make(s.CreateCompany))
}
