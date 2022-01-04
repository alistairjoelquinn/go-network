package database

import (
	"context"
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
			return nil, err
		}
		recentUsers = append(recentUsers, user)
	}

	return &recentUsers, nil
}

func (m DB) UserSearch(q string, id string) (*[]model.RecentUsers, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		SELECT first, last, id, image FROM users
		WHERE first ILIKE $1 AND users.id <> $2
		OR last ILIKE $1 AND users.id <> $2
		OR concat(first, ' ', last) ILIKE $1 AND users.id <> $2
	`

	var recentUsers []model.RecentUsers

	rows, err := m.db.QueryContext(ctx, query, q+"%", id)
	if err != nil {
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
			return nil, err
		}
		recentUsers = append(recentUsers, user)
	}

	return &recentUsers, nil
}

func (m DB) GetOtherUserData(id string) (*model.OtherUser, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "SELECT first, last, image, bio, id as userId FROM users WHERE id = $1"

	var otherUser model.OtherUser

	err := m.db.QueryRowContext(ctx, query, id).Scan(
		&otherUser.First,
		&otherUser.Last,
		&otherUser.Image,
		&otherUser.Bio,
		&otherUser.UserId,
	)
	if err != nil {
		return nil, err
	}

	return &otherUser, err
}
