package integration_test

import (
	"context"
	"log/slog"
	"os"
	"path/filepath"
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
		err := db.CreateCompany(company, user)
		assert.Nil(t, err)
	})

	t.Run("Should error when creating a company with an existing RUC", func(t *testing.T) {
		err := db.CreateCompany(company, user)
		assert.NotNil(t, err)
		assert.Equal(t, "ERROR: duplicate key value violates unique constraint \"company_ruc_key\" (SQLSTATE 23505)", err.Error())
	})

	t.Run("Should error when creating a company with an existing Name", func(t *testing.T) {
		company.RUC = "987654321"
		err := db.CreateCompany(company, user)
		assert.NotNil(t, err)
		assert.Equal(t, "ERROR: duplicate key value violates unique constraint \"company_name_key\" (SQLSTATE 23505)", err.Error())
	})
}
