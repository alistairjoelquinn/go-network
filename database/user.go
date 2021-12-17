package database

import (
	"context"
	"time"
)

func (m DB) UpdateUserBio(id string, bio string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		UPDATE users SET bio = $2
		WHERE id = $1
	`

	_, err := m.db.ExecContext(ctx, query, id, bio)
	if err != nil {
		return err
	}

	return nil
}
