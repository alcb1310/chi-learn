package database

import (
	"database/sql"
	"log/slog"

	"github.com/google/uuid"
)

type CreateUser struct {
	User
	Password string
}

type User struct {
	ID        uuid.UUID
	Email     string
	Name      string
	CompanyID uuid.UUID
}

func addUser(tx *sql.Tx, c *CreateUser) error {
	query := `INSERT INTO "user" (email, name, password, company_id) VALUES ($1, $2, $3, $4) returning id`

	if err := tx.QueryRow(query, c.Email, c.Name, c.Password, c.CompanyID).Scan(&c.ID); err != nil {
		slog.Error("Error creating user", "error", err)
		return err
	}

	return nil
}
