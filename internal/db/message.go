package db

import (
	"context"
	"log"
	"slices"
	"strings"
	"time"

	"github.com/emyasa/yasaworks/internal/ctxkeys"
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

type MessageCursor struct {
	CreatedAt time.Time
	FetchNext bool
}

type Conversation struct {
	ClientFingerprint string
	LatestMessage string
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

func (db *DB) ListMessages(ctx context.Context, clientFingerprint string, limit int, cursor *MessageCursor) ([]Message, error) {
	ctx, span := tracer.Start(ctx, "ListMessages")
	defer span.End()

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	query := `
	SELECT client_fingerprint, sender_type, content, created_at
	FROM messages
	WHERE client_fingerprint = ?`

	args := []any{clientFingerprint}

	if cursor != nil {
		cursorQuery := "AND created_at <= ?"
		if cursor.FetchNext {
			cursorQuery = "AND created_at >= ?"
		}
		query += cursorQuery

		args = append(args,
			cursor.CreatedAt.Format("2006-01-02 15:04:05"),
		)
	}

	cursorQuery := "ORDER BY created_at DESC, id DESC "
	if cursor != nil && cursor.FetchNext {
		cursorQuery = "ORDER BY created_at ASC, id DESC "
	}

	query += cursorQuery
	query += "LIMIT ?"

	args = append(args, limit)

	rows, err := db.handle.QueryContext(ctx, query, args...)
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

	if cursor == nil || !cursor.FetchNext {
		slices.Reverse(messages)
	}

	return messages, nil
}

func (db *DB) ListMessagesByFPs(ctx context.Context, clientFingerprints [] string) []Message {
	if len(clientFingerprints) == 0 {
		log.Fatal("[db][messages] ListMessagesByFPs: clientFingerprints must not be empty")
	}

	ctx, span := tracer.Start(ctx, "ListMessagesMap")
	defer span.End()

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	isAdmin, ok := ctx.Value(ctxkeys.IsAdmin).(bool)
	if !ok || !isAdmin {
		log.Fatal("Non admin clients should not be able to list messages map")
	}

	placeholders := make([]string, len(clientFingerprints))
	cfpsArgs := make([]any, len(clientFingerprints))
	for i, fp := range clientFingerprints {
		placeholders[i] = "?"
		cfpsArgs[i] = fp
	}

	query := `
	SELECT client_fingerprint, sender_type, content, created_at
	FROM (
		SELECT client_fingerprint, sender_type, content, created_at,
		ROW_NUMBER () OVER (
			PARTITION BY client_fingerprint
		) AS rn
		FROM messages
		WHERE client_fingerprint IN (` + strings.Join(placeholders, ",") + `)
		ORDER BY created_at DESC
	) WHERE rn <= 30`

	rows, err := db.handle.QueryContext(ctx, query, cfpsArgs...)
	if err != nil {
		log.Fatalf("ListMessagesMap error %s", err)
	}

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

		t, _ := time.Parse("2006-01-02 15:04:05", createdAt)
		message.CreatedAt = t
		messages = append(messages, message)
	}

	slices.Reverse(messages)
	return messages
}

func (db *DB) ListConversations(ctx context.Context) []Conversation {
	ctx, span := tracer.Start(ctx, "ListConversations")
	defer span.End()

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	isAdmin, ok := ctx.Value(ctxkeys.IsAdmin).(bool)
	if !ok || !isAdmin {
		log.Fatal("Non admin clients should not be able to list conversations")
	}

	query := "SELECT client_fingerprint, content " +
	"FROM (" +
		"SELECT client_fingerprint, content, created_at, " +
		"ROW_NUMBER () OVER (" +
			"PARTITION BY client_fingerprint " +
			"ORDER BY created_at DESC " +
		") AS rn " +
		"FROM messages " +
	") WHERE rn = 1 " +
	"ORDER BY created_at DESC"

	rows, err := db.handle.QueryContext(ctx, query)
	if err != nil {
		log.Fatalf("ListConversations error %s", err)
	}

	conversations := []Conversation{}
	for rows.Next() {
		conversation := Conversation{}
		rows.Scan(&conversation.ClientFingerprint, &conversation.LatestMessage)

		conversations = append(conversations, conversation)
	}

	return conversations
}

