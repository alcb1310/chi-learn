package server

import (
	"net/http"

	"chi-learn/externals/views/home"
)

func (s *Service) Home(w http.ResponseWriter, r *http.Request) error {
	return render(w, r, home.Index())
}
