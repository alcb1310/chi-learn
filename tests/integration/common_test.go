package integration_test

import (
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"chi-learn/internals/database"
	"chi-learn/internals/server"
)

func moun(db database.Service) *server.Service {
	s := server.New(slog.Default(), db)
	s.MountHandlers()
	return s
}

func executeRequest(t *testing.T, s *server.Service, method, url string, body io.Reader) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	req, err := http.NewRequest(method, url, body)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	assert.NoError(t, err)
	s.Router.ServeHTTP(rr, req)

	return rr
}
