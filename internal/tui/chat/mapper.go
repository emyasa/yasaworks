package chat

import "github.com/emyasa/yasaworks/internal/db"

func mapMessages(dbMessages []db.Message) []message {
	messages := []message{}
	for _, m := range dbMessages {
		messages = append(messages, message{
			content: m.Content,
			timestamp: m.CreatedAt,
			isFromSender: m.SenderType == db.SenderClient,
		})
	}

	return messages
}

