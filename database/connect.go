package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/alistairjoelquinn/go-network/util"
)

func ConnectToDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", util.Env("DB_URL"))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
