package admin

import (
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/emyasa/yasaworks/internal/registry"
)

func (m Model) updateChats(messageEvent registry.MessageEvent, isSender bool) {
	message := message {
		content: messageEvent.Message,
		timestamp: time.Now(),
		isFromSender: isSender,
	}

	m.messages[messageEvent.Fingerprint] = append(m.messages[messageEvent.Fingerprint], message)
}

func (m Model) chatPanelView() string {
	if len(m.conversations) == 0 {
		return ""
	}

	selectedConversation := m.conversations[m.selectedConversationIndex]

	sb := strings.Builder{}
	messages := m.messages[selectedConversation.fingerprint]
	for _, message := range messages {
		bubbleStyle := m.theme.ReceiverBubbleStyle()
		position := lipgloss.Left

		if message.isFromSender {
			bubbleStyle = m.theme.SenderBubbleStyle()
			position = lipgloss.Right
		}

		timestamp := message.timestamp.Format("15:04")
		timestampView := m.theme.TimestampStyle().Render(timestamp)

		messageView := lipgloss.Place(m.ViewportWidth - ConversationsViewWidth - 1, 1,
			position, lipgloss.Bottom,
			bubbleStyle.Render(message.content) + timestampView)

		sb.WriteString(messageView + "\n")
	}

	inputView := m.theme.Base().
		MarginLeft(1).
		Render(m.input.View())

	modeString := ""
	if m.Mode == Insert {
		modeString = "-- INSERT --"
	}

	statusLine := m.theme.Base().
		MarginLeft(1).
		Render(modeString)

	child := lipgloss.JoinVertical(lipgloss.Left, sb.String(), inputView, statusLine)

	return lipgloss.Place(m.ViewportWidth - ConversationsViewWidth - 1, m.ViewportHeight, lipgloss.Left, lipgloss.Bottom, child)
}
