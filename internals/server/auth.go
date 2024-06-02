package server

import (
	"net/http"
)

func (s *Service) Login(w http.ResponseWriter, r *http.Request) error {
	r.ParseForm()

	email := r.Form.Get("email")
	if err := validateEmail(email, true); err != nil {
		return err
	}

	password := r.Form.Get("password")
	if err := validatePassword(password, true); err != nil {
		return err
	}

	if err := s.DB.Login(email, password); err != nil {
		return err
	}

	return nil
}
