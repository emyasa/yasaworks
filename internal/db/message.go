package db

import (
	"context"
	"time"

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

type Message struct {
	ClientFingerprint string
	SenderType SenderType
	Content string
	CreatedAt time.Time
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

func (db *DB) ListMessages(ctx context.Context, clientFingerprint string) ([]Message, error) {
	ctx, span := tracer.Start(ctx, "ListMessages")
	defer span.End()

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	rows, err := db.handle.QueryContext(ctx, "SELECT client_fingerprint, sender_type, content, created_at FROM messages WHERE client_fingerprint = ?", clientFingerprint)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	messages := []Message{}
	for rows.Next() {
		message := Message{}
		var createdAt string

		rows.Scan(
			&message.ClientFingerprint,
			&message.SenderType,
			&message.Content,
			&createdAt,
		)

		t, err := time.Parse("2006-01-02 15:04:05", createdAt)
		if err != nil {
			return nil, err
		}

		message.CreatedAt = t
		messages = append(messages, message)
	}

	return messages, nil
}

