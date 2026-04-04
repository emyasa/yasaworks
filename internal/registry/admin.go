package registry

import (
	"context"

	"github.com/google/uuid"
)

var adminConnRegistry = map[string]*Connection{}

func RegisterAdminConnection(ctx context.Context) *Connection {
	adminConn := &Connection{
		id: uuid.NewString(),
		messageChannel: make(chan MessageEvent),
	}

	adminConnRegistry[adminConn.id] = adminConn
	go func() {
		<- ctx.Done()
		delete(adminConnRegistry, adminConn.id)
		close(adminConn.messageChannel)
	}()

	return adminConn
}

func HandleClientMessage(conn *Connection, message string) {
	messageEvent := MessageEvent{
		Fingerprint: conn.fingerprint,
		Message: message,
	}

	for _, conn := range adminConnRegistry {
		select {
		case conn.messageChannel <- messageEvent:
		default:
		}
	}
}


