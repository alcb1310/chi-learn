package server

import "net/http"

func (s *Service) Home(w http.ResponseWriter, r *http.Request) error {
	s.logger.Info("Home")

	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Home Page"))

	return err
}
