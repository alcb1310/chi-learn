package database

import (
	"log/slog"
	"os"
	"strings"
)

func (d *service) Migration() error {
	file, err := os.ReadFile("./internals/database/schema.sql")
	if err != nil {
		slog.Error("Error loading sql file", "error", err)
		return err
	}

	tx, err := d.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	requests := strings.Split(string(file), ";")
	slog.Debug("Migrating database", "requests", requests)
	for _, request := range requests {
		if _, err := tx.Exec(request); err != nil {
			slog.Error("Error migrating database", "error", err)
			return err
		}
	}
	tx.Commit()
	return nil
}
