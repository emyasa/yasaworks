package chat

import (
	"log"
	"time"

	"github.com/emyasa/yasaworks/internal/db"
)

func (m Model) canScrollUp() bool {
	return m.Mode == Normal && m.messagesCursorIndex - messagesWindowSize >= 0
}

func (m *Model) scrollUp() {
	m.messagesCursorIndex--

	if m.hasReachedStart || m.messagesCursorIndex >= messagesBufferSize/2 {
		return
	}

	m.bufferPrevious()
}

func (m *Model) bufferPrevious() {
	message := m.messages[m.messagesCursorIndex + messagesWindowSize]
	m.messages = m.fetchMessages(message.timestamp, false)

	m.messagesCursorIndex += messagesWindowSize
	if len(m.messages) < messagesBufferSize {
		m.messagesCursorIndex = len(m.messages) - messagesWindowSize - 1
		m.hasReachedStart = true
	}

	m.hasReachedEnd = false
}

func (m Model) fetchMessages(timestamp time.Time, next bool) []message {
	msgs, err := m.db.ListMessages(
		m.ctx,
		m.conn.Fingerprint,
		&db.MessageCursor{
			CreatedAt: timestamp,
			FetchNext: next,
		},
	)

	if err != nil {
		log.Fatalf("error: %s", err)
	}

	return mapMessages(msgs)
}

