package server_test

import (
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"

	"chi-learn/internals/database"
	"chi-learn/mocks"
)

type companyTestCase struct {
	name     string
	c        *database.Company
	u        *database.CreateUser
	form     url.Values
	status   int
	err      error
	response string
}

var validCompanyData = []companyTestCase{
	{
		name: "Should create a company",
		c: &database.Company{
			RUC:       "123456789",
			Name:      "Andres",
			IsActive:  true,
			Employees: 3,
		},
		u: &database.CreateUser{
			Password: "test",
			User: database.User{
				Email: "test@test.com",
				Name:  "test",
			},
		},
		form: url.Values{
			"ruc":       {"123456789"},
			"name":      {"Andres"},
			"employees": {"3"},
			"email":     {"test@test.com"},
			"username":  {"test"},
			"password":  {"test"},
		},
		status:   http.StatusCreated,
		err:      nil,
		response: "",
	},
	{
		name: "Should check for conflicts",
		c: &database.Company{
			RUC:       "123456789",
			Name:      "Andres",
			IsActive:  true,
			Employees: 3,
		},
		u: &database.CreateUser{
			Password: "test",
			User: database.User{
				Email: "test@test.com",
				Name:  "test",
			},
		},
		form: url.Values{
			"ruc":       {"123456789"},
			"name":      {"Andres"},
			"employees": {"3"},
			"email":     {"test@test.com"},
			"username":  {"test"},
			"password":  {"test"},
		},
		status:   http.StatusConflict,
		err:      &pgconn.PgError{Code: "23505", Message: "duplicate key value violates unique constraint \"company_ruc_key\""},
		response: "duplicate key value violates unique constraint \"company_ruc_key\"",
	},
}

var invalidCompanyData = []companyTestCase{
	{
		name:     "Should have a RUC",
		c:        &database.Company{},
		u:        &database.CreateUser{},
		form:     url.Values{},
		status:   http.StatusBadRequest,
		err:      nil,
		response: "RUC es requerido",
	},
	{
		name: "Should have a name",
		c:    &database.Company{},
		u:    &database.CreateUser{},
		form: url.Values{
			"ruc": {"123456789"},
		},
		status:   http.StatusBadRequest,
		err:      nil,
		response: "Nombre es requerido",
	},
	{
		name: "Employees can not be zero",
		c:    &database.Company{},
		u:    &database.CreateUser{},
		form: url.Values{
			"ruc":       {"123456789"},
			"name":      {"Andres"},
			"employees": {"0"},
		},
		status:   http.StatusBadRequest,
		err:      nil,
		response: "Empleados no pueden ser cero",
	},
	{
		name: "Employees can not be negative",
		c:    &database.Company{},
		u:    &database.CreateUser{},
		form: url.Values{
			"ruc":       {"123456789"},
			"name":      {"Andres"},
			"employees": {"-1"},
		},
		status:   http.StatusBadRequest,
		err:      nil,
		response: "Empleados no pueden ser negativos",
	},
	{
		name: "Employees can not be a string",
		c:    &database.Company{},
		u:    &database.CreateUser{},
		form: url.Values{
			"ruc":       {"123456789"},
			"name":      {"Andres"},
			"employees": {"test"},
		},
		status:   http.StatusBadRequest,
		err:      nil,
		response: "Empleados no pueden ser cero",
	},
	{
		name: "Email can not be empty",
		c:    &database.Company{},
		u:    &database.CreateUser{},
		form: url.Values{
			"ruc":       {"123456789"},
			"name":      {"Andres"},
			"employees": {"3"},
		},
		status:   http.StatusBadRequest,
		err:      nil,
		response: "Email es requerido",
	},
	{
		name: "Should have a valid email",
		c:    &database.Company{},
		u:    &database.CreateUser{},
		form: url.Values{
			"ruc":   {"123456789"},
			"name":  {"Andres"},
			"email": {"invalid"},
		},
		status:   http.StatusBadRequest,
		err:      nil,
		response: "Email no es válido",
	},
	{
		name: "Password can not be empty",
		c:    &database.Company{},
		u:    &database.CreateUser{},
		form: url.Values{
			"ruc":       {"123456789"},
			"name":      {"Andres"},
			"employees": {"3"},
			"email":     {"test@test.com"},
		},
		status:   http.StatusBadRequest,
		err:      nil,
		response: "Contraseña es requerido",
	},
	{
		name: "Should have a name for the user",
		c:    &database.Company{},
		u:    &database.CreateUser{},
		form: url.Values{
			"ruc":       {"123456789"},
			"name":      {"Andres"},
			"employees": {"3"},
			"email":     {"test@test.com"},
			"password":  {"test"},
		},
		status:   http.StatusBadRequest,
		err:      nil,
		response: "Nombre del usuario es requerido",
	},
}

func TestCreateCompany(t *testing.T) {
	t.Run("Valid data", func(t *testing.T) {
		for _, tc := range validCompanyData {
			t.Run(tc.name, func(t *testing.T) {
				buf := strings.NewReader(tc.form.Encode())
				db := mocks.NewService(t)
				db.EXPECT().CreateCompany(tc.c, tc.u).Return(tc.err)
				s := mount(db)

				rr := executeRequest(t, s, "POST", "/register", buf)
				assert.Equal(t, tc.status, rr.Code)
			})
		}
	})

	t.Run("Invalid data", func(t *testing.T) {
		for _, tc := range invalidCompanyData {
			t.Run(tc.name, func(t *testing.T) {
				buf := strings.NewReader(tc.form.Encode())
				db := mocks.NewService(t)
				db.AssertNotCalled(t, "CreateCompany", tc.c, tc.u)
				s := mount(db)

				rr := executeRequest(t, s, "POST", "/register", buf)
				assert.Equal(t, tc.status, rr.Code)
				assert.Contains(t, rr.Body.String(), tc.response)
			})
		}
	})
}
