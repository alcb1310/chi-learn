package database

import (
	"log/slog"

	"github.com/google/uuid"
)

type Company struct {
	ID        uuid.UUID
	RUC       string
	Name      string
	Employees uint
	IsActive  bool
}

func (s *service) CreateCompany(c *Company, u *CreateUser) error {
	tx, err := s.DB.Begin()
	if err != nil {
		slog.Error("Error creating transaction", "error", err)
		return err
	}
	defer tx.Rollback() // The rollback will happen only if we exit the function without commiting

	query := `INSERT INTO company (ruc, name, employees, is_active) VALUES ($1, $2, $3, $4) returning id`

	if err := tx.QueryRow(query, c.RUC, c.Name, c.Employees, c.IsActive).Scan(&c.ID); err != nil {
		slog.Error("Error creating company", "error", err)
		return err
	}
	u.CompanyID = c.ID

	if err := addUser(tx, u); err != nil {
		return err
	}

	tx.Commit()
	return nil
}
