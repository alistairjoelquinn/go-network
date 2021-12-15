package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/alistairjoelquinn/go-network/model"
)

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

	var id string

	err := m.db.QueryRowContext(ctx, query, first, last, email, hashedPass).Scan(
		&id,
	)
	if err != nil {
		return "", err
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

	err := m.db.QueryRowContext(ctx, query, email).Scan(
		&response.HashedPassword,
		&response.ID,
	)
	if err != nil {
		return &response, err
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

	err := m.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.First,
		&user.Last,
		&user.Image,
		&user.Bio,
	)
	if err != nil {
		return &user, err
	}

	return &user, nil
}

func (m DB) AddNewUserImage(userId string, imageUrl string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "UPDATE users SET image = $2 WHERE id = $1"

	_, err := m.db.ExecContext(ctx, query, userId, imageUrl)
	if err != nil {
		return err
	}

	return nil
}
