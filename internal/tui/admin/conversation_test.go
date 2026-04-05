package admin

import (
	"reflect"
	"testing"

	"github.com/emyasa/yasaworks/internal/registry"
)

func TestUpdateConversations(t *testing.T) {
	events := []registry.MessageEvent{
		{Fingerprint: "A", Message: "msg1"},
		{Fingerprint: "B", Message: "msg2"},
	}

	expected := []conversation{
		{fingerprint: "B", message: "msg2"},
		{fingerprint: "A", message: "msg1"},
	}

	for _, e := range events {
		updateConversations(e)
	}

	if !reflect.DeepEqual(expected, conversations) {
		t.Errorf("expected: %+v, got %+v", expected, conversations)
	}
}

