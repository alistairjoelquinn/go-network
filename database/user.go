package database

import (
	"context"
	"log"
	"time"

	"github.com/alistairjoelquinn/go-network/model"
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

func (m DB) RecentUserSearch(id string) (*[]model.RecentUsers, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		SELECT first, last, id, image FROM users
		WHERE id <> $1
		ORDER BY id DESC
		LIMIT 3
	`

	var recentUsers []model.RecentUsers

	rows, err := m.db.QueryContext(ctx, query, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user model.RecentUsers
		if err := rows.Scan(
			&user.First,
			&user.Last,
			&user.ID,
			&user.Image,
		); err != nil {
			log.Println(err)
			return nil, err
		}
		recentUsers = append(recentUsers, user)
	}

	return &recentUsers, nil
}
