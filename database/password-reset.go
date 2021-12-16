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
