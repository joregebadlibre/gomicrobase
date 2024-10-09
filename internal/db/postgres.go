package db

import (
	"database/sql"
	"gomicrobase/config"

	_ "github.com/lib/pq"
)

func NewPostgresConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", config.GetPostgresDSN())
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
