// Package registry
package registry

type MessageEvent struct {
	Fingerprint string
	Message string
}

type Connection struct {
	id string
	fingerprint string
	messageChannel chan MessageEvent
}

func (c Connection) FetchMessage() MessageEvent {
	return <- c.messageChannel
}

