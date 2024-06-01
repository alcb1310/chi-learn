package server

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"

	"chi-learn/externals/views/bca"
)

type BCAService struct {
	Service

	Router chi.Router
}

func (s *BCAService) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Debug("AuthMiddleware")

		next.ServeHTTP(w, r)
	})
}

func (s *BCAService) BCAHome(w http.ResponseWriter, r *http.Request) error {
	return render(w, r, bca.Index())
}
