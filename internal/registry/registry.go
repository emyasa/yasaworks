// Package registry
package registry

type Connection struct {
	id string
	messageChannel chan string
}

func (c Connection) FetchMessage() string {
	return <- c.messageChannel
}

