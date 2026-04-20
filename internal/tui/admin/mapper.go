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

func mapMessages(dbMessages []db.Message) map[string][]message {
	messagesMap := map[string][]message{}
	for _, m := range dbMessages {
		message :=  message{
			content: m.Content,
			timestamp: m.CreatedAt,
			isFromSender: m.SenderType == db.SenderAdmin,
		}

		messagesMap[m.ClientFingerprint] = append(messagesMap[m.ClientFingerprint], message)
	}

	return messagesMap
}

