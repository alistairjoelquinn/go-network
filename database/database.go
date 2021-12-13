package database

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/alistairjoelquinn/go-network/model"
)

// DB gorm connector
type Login struct {
	hashedPassword string
	id             string
}
type DB struct {
	db *sql.DB
}

var DBModel DB

func ModelInit(db *sql.DB) DB {
	DBModel.db = db

	return DBModel
}

func (m DB) AddNewUser(first string, last string, email string, hashedPass string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		INSERT INTO users (first, last, email, password)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`
	rows, err := m.db.QueryContext(ctx, query, first, last, email, hashedPass)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var id string
	for rows.Next() {
		if err := rows.Scan(
			&id,
		); err != nil {
			return "", err
		}
	}

	return id, nil
}

func (m DB) GetUserPasswordFromEmail(email string) (*model.LogUserIn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		SELECT password, id FROM users 
		WHERE email = $1
	`

	var response model.LogUserIn

	rows, err := m.db.QueryContext(ctx, query, email)
	if err != nil {
		return &response, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&response.HashedPassword,
			&response.ID,
		); err != nil {
			log.Println("response", response)
			return &response, err
		}
	}

	return &response, nil
}

func (m DB) GetUserData(id string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		SELECT id, first, last, image, bio FROM users 
		WHERE id = $1
	`

	var user model.User

	rows, err := m.db.QueryContext(ctx, query, id)
	if err != nil {
		return &user, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&user.ID,
			&user.First,
			&user.Last,
			&user.Image,
			&user.Bio,
		); err != nil {
			log.Println("response", user)
			return &user, err
		}
	}

	return &user, nil
}
