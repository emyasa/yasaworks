package admin

import (
	"reflect"
	"testing"

	"github.com/emyasa/yasaworks/internal/registry"
)

func resetState() {
	conversations = []conversation{}
	conversationsIndex = map[string]int{}
}

func TestUpdateConversations(t *testing.T) {
	tests := []struct{
		name     string
		events   []registry.MessageEvent
		expected []conversation
	}{
		{
			name: "single insert",
			events: []registry.MessageEvent{
				{Fingerprint: "A", Message: "msg1"},
			},
			expected: []conversation{
				{fingerprint: "A", message: "msg1"},
			},
		},
		{
			name: "multiple inserts",
			events: []registry.MessageEvent{
				{Fingerprint: "A", Message: "msg1"},
				{Fingerprint: "B", Message: "msg2"},
			},
			expected: []conversation{
				{fingerprint: "B", message: "msg2"},
				{fingerprint: "A", message: "msg1"},
			},
		},
		{
			name: "update moves to front",
			events: []registry.MessageEvent{
				{Fingerprint: "A", Message: "msg1"},
				{Fingerprint: "B", Message: "msg2"},
				{Fingerprint: "A", Message: "msg3"},
			},
			expected: []conversation{
				{fingerprint: "A", message: "msg3"},
				{fingerprint: "B", message: "msg2"},
			},
		},
		{
			name: "multiple updates",
			events: []registry.MessageEvent{
				{Fingerprint: "A", Message: "msg1"},
				{Fingerprint: "B", Message: "msg2"},
				{Fingerprint: "C", Message: "msg3"},
				{Fingerprint: "B", Message: "msg4"},
			},
			expected: []conversation{
				{fingerprint: "B", message: "msg4"},
				{fingerprint: "C", message: "msg3"},
				{fingerprint: "A", message: "msg1"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetState()

			for _, e := range tt.events {
				updateConversations(e)
			}

			if !reflect.DeepEqual(tt.expected, conversations) {
				t.Errorf("expected: %+v, got %+v", tt.expected, conversations)
			}

			for i, c := range conversations {
				if idx, ok := conversationsIndex[c.fingerprint]; !ok || i != idx {
					if !ok {
						t.Errorf("index map not found for %s", c.fingerprint)
						continue
					}

					t.Errorf("index map inconsistency for %s: expected %d, got %d", c.fingerprint, i, idx)
				}
			}
		})
	}
}

