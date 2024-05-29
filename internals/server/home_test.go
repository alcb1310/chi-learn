package server_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"chi-learn/mocks"
)

func TestHome(t *testing.T) {
	t.Run("Should return the home page", func(t *testing.T) {
		db := mocks.NewService(t)
		s := mount(db)
		rr := executeRequest(t, s, "GET", "/", nil)
		assert.Equal(t, http.StatusOK, rr.Code)
		require.Contains(t, rr.Body.String(), "Hello, World!")
	})
}
