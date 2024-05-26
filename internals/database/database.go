package database

import (
	"database/sql"
	"errors"
	"log/slog"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
)

var (
	connStr = os.Getenv("DATABASE_URL")
)

type Service interface {
	Migration() error
}

type service struct {
	DB *sql.DB
}

func Connect() (Service, error) {
	if connStr == "" {
		slog.Error("DATABASE_URL not set")
		return nil, errors.New("DATABASE_URL not set")
	}
	s := &service{}
	slog.Debug("Connecting to database")

	db, err := sql.Open("pgx", connStr)
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
