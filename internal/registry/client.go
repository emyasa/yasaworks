package registry

import (
	"context"

	"github.com/google/uuid"
)

var clientConnRegistry = map[string]*Connection{}

func RegisterClientConnection(ctx context.Context) *Connection {
	clientConn := &Connection{
		id: uuid.NewString(),
		fingerprint: ctx.Value("fingerprint").(string),
		messageChannel: make(chan MessageEvent),
	}

	clientConnRegistry[clientConn.id] = clientConn
	go func() {
		<- ctx.Done()
		delete(clientConnRegistry, clientConn.id)
		close(clientConn.messageChannel)
	}()

	return clientConn
}

