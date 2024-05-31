package server

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5/pgconn"

	"chi-learn/externals/views/register"
	"chi-learn/internals/database"
	"chi-learn/internals/errs"
)

func (s *Service) Register(w http.ResponseWriter, r *http.Request) error {
	return render(w, r, register.Index())
}

func (s *Service) CreateCompany(w http.ResponseWriter, r *http.Request) error {
	company := &database.Company{
		IsActive: true,
	}
	user := &database.CreateUser{}

	r.ParseForm()
	company.RUC = r.Form.Get("ruc")
	if company.RUC == "" {
		return errs.New(http.StatusBadRequest, "RUC es requerido")
	}

	company.Name = r.Form.Get("name")
	if company.Name == "" {
		return errs.New(http.StatusBadRequest, "Nombre es requerido")
	}

	var empNum uint
	emp := r.Form.Get("employees")
	if emp == "" {
		empNum = 1
	} else {
		num, err := strconv.Atoi(emp)
		if err != nil {
			return errs.New(http.StatusBadRequest, "Empleados tiene que ser un número válido")
		}
		if num == 0 {
			return errs.New(http.StatusBadRequest, "Empleados no pueden ser cero")
		}
		if num < 0 {
			return errs.New(http.StatusBadRequest, "Empleados no pueden ser negativos")
		}

		empNum = uint(num)
	}
	company.Employees = empNum

	user.Email = r.Form.Get("email")
	if err := validateEmail(user.Email, true); err != nil {
		return err
	}

	user.Password = r.Form.Get("password")
	if err := validatePassword(user.Password, true); err != nil {
		return err
	}

	user.Name = r.Form.Get("username")
	if user.Name == "" {
		return errs.New(http.StatusBadRequest, "Nombre del usuario es requerido")
	}

	if err := s.DB.CreateCompany(company, user); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			slog.Error("Error creating company", "error", err, "method", r.Method, "url", r.URL.Path, "code", pgErr.Code, "error", pgErr)
			if pgErr.Code == "23505" {
				return errs.New(http.StatusConflict, pgErr.Message)
			}
		}
		return err
	}

	w.WriteHeader(http.StatusCreated)
	return nil
}
