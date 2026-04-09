package admin

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/emyasa/yasaworks/internal/registry"
)

func (m Model) updateChats(messageEvent registry.MessageEvent, isSender bool) {
	messagePosition := lipgloss.Left
	if isSender {
		messagePosition = lipgloss.Right
	}

	message := lipgloss.Place(80, 1, messagePosition, lipgloss.Bottom, messageEvent.Message)
	m.messages[messageEvent.Fingerprint] = append(m.messages[messageEvent.Fingerprint], message)
}

func (m Model) chatPanelView() string {
	var messagesView string
	if len(m.conversations) > 0 {
		selectedConversation := m.conversations[m.selectedConversationIndex]
		messages := m.messages[selectedConversation.fingerprint]
		messagesView = strings.Join(messages, "\n")
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

	child := lipgloss.JoinVertical(lipgloss.Left, messagesView, inputView, statusLine)

	return lipgloss.Place(80, 23, lipgloss.Left, lipgloss.Bottom, child)
}

