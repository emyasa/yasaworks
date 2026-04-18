package admin

import "github.com/emyasa/yasaworks/internal/db"

func mapConversations(dbConversations []db.Conversation) ([]conversation, map[string]int) {
	conversations := []conversation{}
	for _, c := range dbConversations {
		conversations = append(conversations, conversation{
			fingerprint: c.ClientFingerprint,
			message: c.LatestMessage,
		})
	}

	conversationsIndex := map[string]int{}
	for i, c := range conversations {
		conversationsIndex[c.fingerprint] = i
	}

	return conversations, conversationsIndex
}

func mapMessages(dbMessages []db.Message) []message {
	messages := []message{}
	for _, m := range dbMessages {
		messages = append(messages, message{
			content: m.Content,
			timestamp: m.CreatedAt,
			isFromSender: m.SenderType == db.SenderAdmin,
		})
	}

	return messages
}

