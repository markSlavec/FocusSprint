package main

import (
	"database/sql"
	"fmt"
)

func applyMigrations(db *sql.DB) error {
	// создаём таблицы, если их нет (SQLite синтаксис)
	mattersTable := `
		CREATE TABLE IF NOT EXISTS matters (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)
	`

	subMattersTable := `
		CREATE TABLE IF NOT EXISTS sub_matters (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			matter_id INTEGER NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (matter_id) REFERENCES matters(id) ON DELETE CASCADE
		)
	`

	if _, err := db.Exec(mattersTable); err != nil {
		return fmt.Errorf("создание таблицы matters: %w", err)
	}

	if _, err := db.Exec(subMattersTable); err != nil {
		return fmt.Errorf("создание таблицы sub_matters: %w", err)
	}

	return nil
}
