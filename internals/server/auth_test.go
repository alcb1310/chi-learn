package server_test

import (
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"chi-learn/mocks"
)

type authTestCase struct {
	name     string
	form     url.Values
	status   int
	err      error
	response string
}

var invalidLoginData = []authTestCase{
	{
		name:     "Email can not be empty",
		form:     url.Values{},
		status:   http.StatusBadRequest,
		err:      nil,
		response: "Email es requerido",
	},
	{
		name: "Should have a valid email",
		form: url.Values{
			"email": {"invalid"},
		},
		status:   http.StatusBadRequest,
		err:      nil,
		response: "Email no es válido",
	},
	{
		name: "Password can not be empty",
		form: url.Values{
			"email": {"test@test.com"},
		},
		status:   http.StatusBadRequest,
		err:      nil,
		response: "Contraseña es requerido",
	},
}

func TestLogin(t *testing.T) {
	for _, c := range invalidLoginData {
		t.Run(c.name, func(t *testing.T) {
			buf := strings.NewReader(c.form.Encode())
			db := mocks.NewService(t)
			db.AssertNotCalled(t, "Login", c.form.Get("email"), c.form.Get("password"))
			s := mount(db)

			rr := executeRequest(t, s, "POST", "/login", buf)
			assert.Equal(t, c.status, rr.Code)
			assert.Contains(t, rr.Body.String(), c.response)
		})
	}
}
