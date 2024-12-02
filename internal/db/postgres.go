package db

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

type PostgresDB struct {
	Client *sql.DB
}

func NewPostgresDB(uri string) (Database, error) {
	db, err := sql.Open("postgres", uri)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresDB{Client: db}, nil
}

func (db *PostgresDB) Disconnect(ctx context.Context) error {
	return db.Client.Close()
}
