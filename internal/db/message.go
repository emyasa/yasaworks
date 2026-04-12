package db

import (
	"context"

	"github.com/emyasa/yasaworks/internal/tracer"
)

type SenderType string
const (
	SenderClient SenderType = "client"
	SenderAdmin SenderType = "admin"
)

type CreateMessageRequest struct {
	ClientFingerprint string
	SenderType SenderType
	Content string
}

func (db *DB) CreateMessage(ctx context.Context, r CreateMessageRequest) error {
	_, span := tracer.Start(ctx, "CreateMessage")
	defer span.End()

	_, err := db.handle.Exec("INSERT INTO messages (client_fingerprint, sender_type, content) VALUES (?, ?, ?)", r.ClientFingerprint, r.SenderType, r.Content)
	if err != nil {
		return err
	}

	return nil
}

