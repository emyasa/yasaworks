package db

import (
	"context"
	"time"
)

type UpsertUserRequest struct {
	fingerprint string
	clientIP string
}

func (db *DB) UpsertUser(ctx context.Context, r UpsertUserRequest) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
}

