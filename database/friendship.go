package database

import (
	"context"
	"log"
	"time"

	"github.com/alistairjoelquinn/go-network/model"
)

func (m DB) GetRequestsFriends(id string) (*[]model.RequestsFriends, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
			SELECT first, last, image, accepted, friendships.id as friendship_id, users.id AS id
			FROM friendships
			JOIN users
			ON (accepted = false AND recipient_id = $1 AND sender_id = users.id)
			OR (accepted = true AND recipient_id = $1 AND sender_id = users.id)
			OR (accepted = true AND sender_id = $1 AND recipient_id = users.id)
		`

	var requestsFriends []model.RequestsFriends

	rows, err := m.db.QueryContext(ctx, query, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user model.RequestsFriends
		if err := rows.Scan(
			&user.First,
			&user.Last,
			&user.Image,
			&user.Accepted,
			&user.FriendshipId,
			&user.ID,
		); err != nil {
			log.Println(err)
			return nil, err
		}
		requestsFriends = append(requestsFriends, user)
	}

	return &requestsFriends, nil
}

func (m DB) FriendshipStatus(userId string, id string) (model.FStatus, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
			SELECT sender_id, recipient_id, accepted, id 
			FROM friendships 
			WHERE (recipient_id = $1 AND sender_id = $2)
			OR (recipient_id = $2 AND sender_id = $1)
		`

	var friendshipStatus model.FStatus

	err := m.db.QueryRowContext(ctx, query, userId, id).Scan(
		&friendshipStatus.SenderId,
		&friendshipStatus.RecipientId,
		&friendshipStatus.Accepted,
		&friendshipStatus.ID,
	)
	if err != nil {
		log.Println(err)
		return friendshipStatus, err
	}

	return friendshipStatus, nil
}
