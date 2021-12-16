package database

import (
	"context"
	"time"
)

func (m DB) CheckEmailForReset(email string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "SELECT id, email FROM users WHERE email = $1"

	var id string

	err := m.db.QueryRowContext(ctx, query, email).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}
