// Package db handles database related operations
package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func New() *sql.DB {
	db, err := sql.Open("sqlite", "data/main.db")
	if err != nil {
		panic(err)
	}

	return db
}

