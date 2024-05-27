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

func TestCreateCompany(t *testing.T) {
	t.Run("Using valid data", func(t *testing.T) {
		t.Run("Should create a company", func(t *testing.T) {
			c := &database.Company{
				RUC:       "123456789",
				Name:      "Andres",
				IsActive:  true,
				Employees: 3,
			}
			u := &database.CreateUser{}
			u.Email = "test@test.com"
			u.Name = "test"
			u.Password = "test"

			form := url.Values{}
			form.Add("ruc", c.RUC)
			form.Add("name", c.Name)
			form.Add("employees", "3")
			form.Add("email", "test@test.com")
			form.Add("username", "test")
			form.Add("password", "test")

			buf := strings.NewReader(form.Encode())

			db := mocks.NewService(t)
			db.EXPECT().CreateCompany(c, u).Return(nil)
			s := mount(db)
			rr := executeRequest(t, s, "POST", "/companies", buf)
			assert.Equal(t, http.StatusCreated, rr.Code)
		})

		t.Run("Should check for conflicts", func(t *testing.T) {
			c := &database.Company{
				RUC:       "123456789",
				Name:      "Andres",
				IsActive:  true,
				Employees: 3,
			}
			u := &database.CreateUser{}
			u.Email = "test@test.com"
			u.Name = "test"
			u.Password = "test"

			form := url.Values{}
			form.Add("ruc", c.RUC)
			form.Add("name", c.Name)
			form.Add("employees", "3")
			form.Add("email", "test@test.com")
			form.Add("username", "test")
			form.Add("password", "test")

			buf := strings.NewReader(form.Encode())

			db := mocks.NewService(t)
			pgErr := &pgconn.PgError{}
			pgErr.Code = "23505"
			pgErr.Message = "duplicate key value violates unique constraint \"company_ruc_key\""
			db.EXPECT().CreateCompany(c, u).Return(pgErr)
			s := mount(db)
			rr := executeRequest(t, s, "POST", "/companies", buf)
			assert.Equal(t, http.StatusConflict, rr.Code)
			assert.Contains(t, rr.Body.String(), pgErr.Message)
		})
	})

	t.Run("Should validate user input", func(t *testing.T) {
		t.Run("Should have a RUC", func(t *testing.T) {
			db := mocks.NewService(t)
			s := mount(db)
			rr := executeRequest(t, s, "POST", "/companies", nil)
			assert.Equal(t, http.StatusBadRequest, rr.Code)
			assert.Contains(t, rr.Body.String(), "RUC es requerido")
		})

		t.Run("Should have a name", func(t *testing.T) {
			form := url.Values{}
			form.Add("ruc", "123456789")
			buf := strings.NewReader(form.Encode())

			db := mocks.NewService(t)
			s := mount(db)
			rr := executeRequest(t, s, "POST", "/companies", buf)
			assert.Equal(t, http.StatusBadRequest, rr.Code)
			assert.Contains(t, rr.Body.String(), "Nombre es requerido")
		})

		t.Run("Employees must be a positive integer", func(t *testing.T) {
			t.Run("Employees can not be zero", func(t *testing.T) {
				form := url.Values{}
				form.Add("ruc", "123456789")
				form.Add("name", "Andres")
				form.Add("employees", "0")
				buf := strings.NewReader(form.Encode())

				db := mocks.NewService(t)
				s := mount(db)
				rr := executeRequest(t, s, "POST", "/companies", buf)
				assert.Equal(t, http.StatusBadRequest, rr.Code)
				assert.Contains(t, rr.Body.String(), "Empleados no pueden ser cero")
			})

			t.Run("Employees can not be negative", func(t *testing.T) {
				form := url.Values{}
				form.Add("ruc", "123456789")
				form.Add("name", "Andres")
				form.Add("employees", "-1")
				buf := strings.NewReader(form.Encode())

				db := mocks.NewService(t)
				s := mount(db)
				rr := executeRequest(t, s, "POST", "/companies", buf)
				assert.Equal(t, http.StatusBadRequest, rr.Code)
				assert.Contains(t, rr.Body.String(), "Empleados no pueden ser negativos")
			})

			t.Run("Employees can not be a string", func(t *testing.T) {
				form := url.Values{}
				form.Add("ruc", "123456789")
				form.Add("name", "Andres")
				form.Add("employees", "1.1")
				buf := strings.NewReader(form.Encode())

				db := mocks.NewService(t)
				s := mount(db)
				rr := executeRequest(t, s, "POST", "/companies", buf)
				assert.Equal(t, http.StatusBadRequest, rr.Code)
				assert.Contains(t, rr.Body.String(), "Empleados tiene que ser un número válido")
			})
		})

		t.Run("Should have a valid email", func(t *testing.T) {
			t.Run("Empty email", func(t *testing.T) {
				form := url.Values{}
				form.Add("ruc", "123456789")
				form.Add("name", "Andres")
				form.Add("employees", "1")
				buf := strings.NewReader(form.Encode())

				db := mocks.NewService(t)
				s := mount(db)
				rr := executeRequest(t, s, "POST", "/companies", buf)
				assert.Equal(t, http.StatusBadRequest, rr.Code)
				assert.Contains(t, rr.Body.String(), "Email es requerido")
			})

			t.Run("Invalid email", func(t *testing.T) {
				form := url.Values{}
				form.Add("ruc", "123456789")
				form.Add("name", "Andres")
				form.Add("email", "test")
				buf := strings.NewReader(form.Encode())

				db := mocks.NewService(t)
				s := mount(db)
				rr := executeRequest(t, s, "POST", "/companies", buf)
				assert.Equal(t, http.StatusBadRequest, rr.Code)
				assert.Contains(t, rr.Body.String(), "Email no es válido")
			})
		})

		t.Run("Should have a valid password", func(t *testing.T) {
			t.Run("Empty password", func(t *testing.T) {
				form := url.Values{}
				form.Add("ruc", "123456789")
				form.Add("name", "Andres")
				form.Add("email", "valid@me.com")
				buf := strings.NewReader(form.Encode())

				db := mocks.NewService(t)
				s := mount(db)
				rr := executeRequest(t, s, "POST", "/companies", buf)
				assert.Equal(t, http.StatusBadRequest, rr.Code)
				assert.Contains(t, rr.Body.String(), "Contraseña es requerido")
			})
		})

		t.Run("Should have a valid name for the user", func(t *testing.T) {
			t.Run("Empty name", func(t *testing.T) {
				form := url.Values{}
				form.Add("ruc", "123456789")
				form.Add("name", "Andres")
				form.Add("email", "valid@me.com")
				form.Add("password", "123456")
				buf := strings.NewReader(form.Encode())

				db := mocks.NewService(t)
				s := mount(db)
				rr := executeRequest(t, s, "POST", "/companies", buf)
				assert.Equal(t, http.StatusBadRequest, rr.Code)
				assert.Contains(t, rr.Body.String(), "Nombre del usuario es requerido")
			})
		})
	})
}
