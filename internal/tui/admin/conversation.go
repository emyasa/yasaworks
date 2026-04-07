package admin

import "github.com/emyasa/yasaworks/internal/registry"

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

