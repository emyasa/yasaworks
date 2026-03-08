package db

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/emyasa/yasaworks/internal/tracer"
)

type UpsertUserRequest struct {
	Fingerprint string
	ClientIP string
}

func (db *DB) UpsertUser(ctx context.Context, r UpsertUserRequest) {
	ctx, span := tracer.Start(ctx, "UpsertUser")
	defer span.End()

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	var userID uint64
	row := db.handle.QueryRowContext(ctx, "SELECT id FROM users WHERE fingerprint = ?", r.Fingerprint)

	err := row.Scan(&userID)
	if err == sql.ErrNoRows {
		result, err := db.handle.Exec("INSERT INTO users (fingerprint) VALUES (?)", r.Fingerprint)
		if err != nil {
			log.Fatal("UpsertUser Error: " + err.Error())
		}

		id, err := result.LastInsertId()
		if err != nil {
			log.Fatal("UpsertUser Error: " + err.Error())
		}

		userID = uint64(id)
	}

	_, err = db.handle.Exec("INSERT INTO login_history (user_id, ip_address) VALUES (?, ?)", userID, r.ClientIP)
	if err != nil {
		log.Fatal("UpsertUser Error: " + err.Error())
	}
}

