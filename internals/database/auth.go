package database

import (
	"net/http"

	"chi-learn/internals/errs"
	"chi-learn/internals/utils"
)

func (s *service) Login(email, password string) error {
	var e, p string

	query := `SELECT email, password FROM "user" WHERE email = $1`
	if err := s.DB.QueryRow(query, email).Scan(&e, &p); err != nil {
		return errs.New(http.StatusUnauthorized, "Credenciales incorrectas")
	}

	if _, err := utils.ComparePassword(p, password); err != nil {
		return errs.New(http.StatusUnauthorized, "Credenciales incorrectas")
	}

	return nil
}
