package database

import (
	"context"
	"time"

	"github.com/alistairjoelquinn/go-network/model"
)

func (m DB) GetRequestsFriends(q string, id string) (*[]model.RequestsFriends, error) {
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

	rows, err := m.db.QueryContext(ctx, query, id)
}
