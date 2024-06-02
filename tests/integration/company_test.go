package integration_test

import (
	"context"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"

	"chi-learn/internals/database"
)

var company = &database.Company{
	RUC:       "123456789",
	Name:      "Andres",
	IsActive:  true,
	Employees: 3,
}

var user = &database.CreateUser{}

func TestCreateCompany(t *testing.T) {
	user.Email = "test@test.com"
	user.Name = "test"
	user.Password = "test"

	ctx := context.Background()

	pgContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("postgres:latest"),
		postgres.WithInitScripts(filepath.Join("..", "..", "internals", "database", "schema.sql")),
		postgres.WithDatabase("test-db"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second),
		),
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		if err := pgContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate pgContainer: %s", err)
		}
	})
	connString, err := pgContainer.ConnectionString(ctx)
	if err != nil {
		t.Fatal(err)
	}
	db, err := database.Connect(connString)
	if err != nil {
		slog.Error("Error connecting to database", "error", err)
		os.Exit(1)
	}

	t.Run("Should create a new company", func(t *testing.T) {
		form := url.Values{}
		form.Add("ruc", "123456789")
		form.Add("name", "Test Company")
		form.Add("employees", "3")
		form.Add("email", "test@test.com")
		form.Add("username", "test")
		form.Add("password", "test")

		buf := strings.NewReader(form.Encode())
		s := moun(db)
		rr := executeRequest(t, s, "POST", "/register", buf)

		assert.Equal(t, http.StatusCreated, rr.Code)
	})

	t.Run("Should error when creating a company with an existing RUC", func(t *testing.T) {
		form := url.Values{}
		form.Add("ruc", "123456789")
		form.Add("name", "Test Company")
		form.Add("employees", "3")
		form.Add("email", "test@test.com")
		form.Add("username", "test")
		form.Add("password", "test")

		buf := strings.NewReader(form.Encode())
		s := moun(db)
		rr := executeRequest(t, s, "POST", "/register", buf)

		assert.Equal(t, http.StatusConflict, rr.Code)
		assert.Equal(t, "duplicate key value violates unique constraint \"company_ruc_key\"\n", rr.Body.String())
	})

	t.Run("Should error when creating a company with an existing Name", func(t *testing.T) {
		form := url.Values{}
		form.Add("ruc", "987654321")
		form.Add("name", "Test Company")
		form.Add("employees", "3")
		form.Add("email", "test@test.com")
		form.Add("username", "test")
		form.Add("password", "test")

		buf := strings.NewReader(form.Encode())
		s := moun(db)
		rr := executeRequest(t, s, "POST", "/register", buf)

		assert.Equal(t, http.StatusConflict, rr.Code)
		assert.Equal(t, "duplicate key value violates unique constraint \"company_name_key\"\n", rr.Body.String())
	})

	t.Run("Should error when creating a company with an existing email", func(t *testing.T) {
		form := url.Values{}
		form.Add("ruc", "987654321")
		form.Add("name", "Another Test Company")
		form.Add("employees", "3")
		form.Add("email", "test@test.com")
		form.Add("username", "test")
		form.Add("password", "test")

		buf := strings.NewReader(form.Encode())
		s := moun(db)
		rr := executeRequest(t, s, "POST", "/register", buf)

		assert.Equal(t, http.StatusConflict, rr.Code)
		assert.Equal(t, "duplicate key value violates unique constraint \"user_email_key\"\n", rr.Body.String())
	})
}
