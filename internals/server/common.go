package server

import (
	"log/slog"
	"net/http"
)

type HTTPFunc func(w http.ResponseWriter, r *http.Request) error

func make(f HTTPFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			slog.Error("Error handling request", "error", err, "method", r.Method, "url", r.URL.Path)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
