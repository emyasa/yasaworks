// Package db handles database related operations
package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type DB struct {
	handle *sql.DB
}

func New() *DB {
	db, err := sql.Open("sqlite", "data/main.db")
	if err != nil {
		panic(err)
	}

	return &DB{handle: db}
}

func (db *DB) Close() {
	db.handle.Close()
}

