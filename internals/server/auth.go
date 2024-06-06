package server

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/jwtauth/v5"
)

func (s *Service) Login(w http.ResponseWriter, r *http.Request) error {
	cookie := &http.Cookie{
		Name:     "bca",
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour),
		MaxAge:   86400,
		Raw:      "",
		HttpOnly: true,
		Secure:   true,
	}

	r.ParseForm()

	email := r.Form.Get("email")
	if err := validateEmail(email, true); err != nil {
		return err
	}

	password := r.Form.Get("password")
	if err := validatePassword(password, true); err != nil {
		return err
	}

	u, err := s.DB.Login(email, password)
	if err != nil {
		return err
	}

	token := jwtauth.New("HS256", []byte(os.Getenv("SECRET")), nil)
	_, tokenString, _ := token.Encode(map[string]interface{}{"id": u.ID, "email": u.Email, "company_id": u.CompanyID, "name": u.Name})
	cookie.Value = tokenString
	slog.Info("Login successful", "cookie", cookie)
	http.SetCookie(w, cookie)
	slog.Debug("Login successful", "token", tokenString)

	return nil
}

func (s *BCAService) Logout(w http.ResponseWriter, r *http.Request) error {
	cookie := &http.Cookie{
		Name:     "bca",
		Value:    "",
		Path:     "/",
		Expires:  time.Now(),
		MaxAge:   -1,
		Raw:      "",
		HttpOnly: true,
		Secure:   true,
	}

	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusFound)
	return nil
}
