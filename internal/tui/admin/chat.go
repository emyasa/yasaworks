package admin

import (
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/emyasa/yasaworks/internal/registry"
)

func (m Model) updateChats(messageEvent registry.MessageEvent, isSender bool) {
	bubbleStyle := m.theme.ReceiverBubbleStyle()
	position := lipgloss.Left

	if isSender {
		bubbleStyle = m.theme.SenderBubbleStyle()
		position = lipgloss.Right
	}

	timestamp := time.Now().Format("15:04")
	timestampView := m.theme.TimestampStyle().Render(timestamp)

	message := lipgloss.Place(80, 1, position, lipgloss.Bottom, bubbleStyle.Render(messageEvent.Message) + timestampView)
	m.messages[messageEvent.Fingerprint] = append(m.messages[messageEvent.Fingerprint], message)
}

func (m Model) chatPanelView() string {
	if len(m.conversations) == 0 {
		return "No messages. Press q to quit."
	}

	selectedConversation := m.conversations[m.selectedConversationIndex]
	messages := m.messages[selectedConversation.fingerprint]
	messagesView := strings.Join(messages, "\n")

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

	child := lipgloss.JoinVertical(lipgloss.Left, messagesView, inputView, statusLine)

	return lipgloss.Place(80, 23, lipgloss.Left, lipgloss.Bottom, child)
}
