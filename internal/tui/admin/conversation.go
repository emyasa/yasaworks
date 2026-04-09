package admin

import (
	"strings"

	"github.com/emyasa/yasaworks/internal/registry"
)

type conversation struct {
	fingerprint string
	message string
}

func (m *Model) updateConversations(messageEvent registry.MessageEvent) {
	fp := messageEvent.Fingerprint
	msg := messageEvent.Message

	if idx, ok := m.conversationsIndex[fp]; ok {
		m.conversations[idx].message = msg
		if idx == 0 {
			return
		}

		updated := m.conversations[idx]
		copy(m.conversations[1:idx+1], m.conversations[0:idx])
		m.conversations[0] = updated

		for i := 0; i <= idx; i++ {
			m.conversationsIndex[m.conversations[i].fingerprint] = i
		}

		return
	}

	m.conversations = append(m.conversations, conversation{})
	copy(m.conversations[1:], m.conversations[:len(m.conversations) - 1])
	m.conversations[0] = conversation{
		fingerprint: fp,
		message: msg,
	}

	m.conversationsIndex[fp] = 0
	for i := 1; i < len(m.conversations); i++ {
		m.conversationsIndex[m.conversations[i].fingerprint] = i
	}
}

func (m Model) conversationsView() string {
	sb := strings.Builder{}
	for i, c := range m.conversations {
		menuItemStyle := m.theme.Base().
			Width(17).
			Padding(0, 0, 0, 1)

		if i == m.selectedConversationIndex {
			menuItemStyle = menuItemStyle.Background(m.theme.Highlight()).
				Foreground(m.theme.Accent()).
				Bold(true)
		}

		sb.WriteString(menuItemStyle.Render(c.fingerprint[:7]))
		sb.WriteString("\n")

		previewMessage := c.message
		if len(previewMessage) > 12 {
			previewMessage = c.message[0:12] + "..."
		}

		sb.WriteString(menuItemStyle.Render(previewMessage))
		sb.WriteString("\n\n")
	}

	return sb.String()
}

