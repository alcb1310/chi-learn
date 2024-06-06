package database

import (
	"net/http"

	"chi-learn/internals/errs"
	"chi-learn/internals/utils"
)

func (s *service) Login(email, password string) (User, error) {
	u := User{}
	var p string

	query := `SELECT id, company_id, email, password, name FROM "user" WHERE email = $1`
	if err := s.DB.QueryRow(query, email).Scan(&u.ID, &u.CompanyID, &u.Email, &p, &u.Name); err != nil {
		return u, errs.New(http.StatusUnauthorized, "Credenciales incorrectas")
	}

	if _, err := utils.ComparePassword(p, password); err != nil {
		return User{}, errs.New(http.StatusUnauthorized, "Credenciales incorrectas")
	}

	return u, nil
}
