package server_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHome(t *testing.T) {
	t.Run("Should return the home page", func(t *testing.T) {
		s := mount()
		rr := executeRequest(t, s, "GET", "/", nil)
		assert.Equal(t, http.StatusOK, rr.Code)
		require.Equal(t, "Home Page", rr.Body.String())
	})
}
