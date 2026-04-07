package admin

import "github.com/emyasa/yasaworks/internal/registry"

type conversation struct {
	fingerprint string
	message string
}

var (
	conversations = make([]conversation, 0)
	conversationsIndex = make(map[string]int)
)

func updateConversations(messageEvent registry.MessageEvent) {
	fp := messageEvent.Fingerprint
	msg := messageEvent.Message

	if idx, ok := conversationsIndex[fp]; ok {
		conversations[idx].message = msg
		if idx == 0 {
			return
		}

		updated := conversations[idx]
		copy(conversations[1:idx+1], conversations[0:idx])
		conversations[0] = updated

		for i := 0; i <= idx; i++ {
			conversationsIndex[conversations[i].fingerprint] = i
		}

		return
	}

	conversations = append(conversations, conversation{})
	copy(conversations[1:], conversations[:len(conversations) - 1])
	conversations[0] = conversation{
		fingerprint: fp,
		message: msg,
	}

	conversationsIndex[fp] = 0
	for i := 1; i < len(conversations); i++ {
		conversationsIndex[conversations[i].fingerprint] = i
	}
}

