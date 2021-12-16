package database

import (
	"context"
	"log"
	"time"
)

func (m DB) CheckEmailForReset(email string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "SELECT id FROM users WHERE email = $1"

	var id string

	err := m.db.QueryRowContext(ctx, query, email).Scan(&id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (m DB) InsertResetCode(code string, email string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "INSERT INTO password_reset_codes (code, email) VALUES ($1, $2)"

	_, err := m.db.ExecContext(ctx, query, code, email)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
