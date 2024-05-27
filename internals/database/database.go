package database

import (
	"database/sql"
	"log/slog"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
)

type Service interface {
	Migration() error

	// Company
	CreateCompany(c *Company, u *CreateUser) error
}

type service struct {
	DB *sql.DB
}

func Connect(dbURL string) (Service, error) {
	s := &service{}
	slog.Debug("Connecting to database")

	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		slog.Error("Error connecting to database", "error", err)
		return nil, err
	}
	s.DB = db

	if err := s.DB.Ping(); err != nil {
		slog.Error("Error pinging database", "error", err)
		return nil, err
	}
	slog.Debug("Connected to database")

	return s, nil
}
