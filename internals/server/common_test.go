package server_test

import (
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"chi-learn/internals/server"
)

func mount() *server.Service {
	s := server.New(slog.Default())
	s.MountHandlers()
	return s
}

func executeRequest(t *testing.T, s *server.Service, method, url string, body io.Reader) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	req, err := http.NewRequest(method, url, body)

	assert.NoError(t, err)
	s.Router.ServeHTTP(rr, req)

	return rr
}
