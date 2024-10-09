package db

import (
	"database/sql"
	"gomicrobase/config"
	"log"

	_ "github.com/lib/pq"
)

func NewPostgresConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", config.GetPostgresDSN())
	if err != nil {
		log.Fatalf("New Connection Open: %v", err.Error())
		return nil, err
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to Ping: %v", err.Error())
		return nil, err
	}
	return db, nil
}
