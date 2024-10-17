package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/Utro-tvar/Storage/internal/db"
	"github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const op = "db.sqlite.New"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS resources(
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL UNIQUE,
		type TEXT NOT NULL,
		path TEXT NOT NULL);
	`)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) Save(path string, resourseType int, name string) error {
	const op = "db.sqlite.Save"

	stmt, err := s.db.Prepare("INSERT INTO resources(name, type, path) VALUES(?, ?, ?)")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec(name, resourseType, path)
	if err != nil {
		if err, ok := err.(sqlite3.Error); ok && err.ExtendedCode == sqlite3.ErrConstraintUnique {
			return fmt.Errorf("%s: %w", op, db.ErrResourceExists)
		}
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
