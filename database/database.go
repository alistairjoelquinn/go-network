package database

import (
	"context"
	"database/sql"
	"time"
)

// DB gorm connector
type DB struct {
	db *sql.DB
}

var DBModel DB

func ModelInit(db *sql.DB) DB {
	DBModel.db = db

	return DBModel
}

func (m DB) AddNewUser(first string, last string, email string, hashedPass string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		INSERT INTO users (first, last, email, password)
		VALUES ($1, $2, $3, $4)
	`
	_, err := m.db.ExecContext(ctx, query, first, last, email, hashedPass)
	if err != nil {
		return err
	}

	return nil
}
