package server

import (
	"net/http"

	"chi-learn/externals/views/home"
)

func (s *Service) Home(w http.ResponseWriter, r *http.Request) error {
	return home.Index().Render(r.Context(), w)
}
