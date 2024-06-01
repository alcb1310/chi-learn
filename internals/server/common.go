package server

import (
	"log/slog"
	"net/http"
	"net/mail"

	"github.com/a-h/templ"

	"chi-learn/internals/errs"
)

type HTTPFunc func(w http.ResponseWriter, r *http.Request) error

func handleErrors(f HTTPFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			re, ok := err.(*errs.Error)
			if ok {
				http.Error(w, err.Error(), re.Code)
				return
			}

			slog.Error("Error handling request", "error", err, "method", r.Method, "url", r.URL.Path)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func validateEmail(email string, required bool) error {
	if email == "" && required {
		return errs.New(http.StatusBadRequest, "Email es requerido")
	}
	_, err := mail.ParseAddress(email)
	if err != nil {
		return errs.New(http.StatusBadRequest, "Email no es válido")
	}

	return nil
}

func validatePassword(password string, required bool) error {
	if password == "" && required {
		return errs.New(http.StatusBadRequest, "Contraseña es requerido")
	}
	return nil
}

func render(w http.ResponseWriter, r *http.Request, c templ.Component) error {
	return c.Render(r.Context(), w)
}
