package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/alistairjoelquinn/go-network/models"
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

func (m DB) AddNewUser(person *models.NewUser) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		INSERT INTO users (first, last, email, password)
		VALUES ($1, $2, $3, $4)
	`
	_, err := m.db.ExecContext(ctx, query, person.First, person.Last, person.Email, person.Password)
	if err != nil {
		return err
	}

	return nil
}
