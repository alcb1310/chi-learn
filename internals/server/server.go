package server

import (
	"log/slog"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Service struct {
	logger *slog.Logger
	Router *chi.Mux
}

func New(logger *slog.Logger) *Service {
	s := &Service{
		logger: logger,
		Router: chi.NewRouter(),
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
}
