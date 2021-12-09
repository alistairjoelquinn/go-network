package database

import (
	"context"
	"database/sql"
	"time"
)

// DB gorm connector
var DB *sql.DB

func AddNewUser(first string, last string, email string, passHash string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		INSERT INTO users (first, last, email, password)
		VALUES ($1, $2, $3, $4)
		RETURNING first, last, id
	`

	queryParams := []string{first, last, email, passHash}

	_, err := DB.ExecContext(ctx, query, queryParams)

	if err != nil {
		return err
	}

	return nil
}
