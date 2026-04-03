package registry

import (
	"context"

	"github.com/google/uuid"
)

var adminConnRegistry = map[string]*Connection{}

func RegisterAdminConnection(ctx context.Context) *Connection {
	adminConn := &Connection{id: uuid.NewString(), messageChannel: make(chan string)}
	adminConnRegistry[adminConn.id] = adminConn
	go func() {
		<- ctx.Done()
		delete(adminConnRegistry, adminConn.id)
		close(adminConn.messageChannel)
	}()

	return adminConn
}

func HandleClientMessage(message string) {
	for _, conn := range adminConnRegistry {
		select {
		case conn.messageChannel <- message:
		default:
		}
	}
}


