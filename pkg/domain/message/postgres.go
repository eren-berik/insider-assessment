package message

import (
	"context"
)

type (
	PostgresService interface {
		Message(ctx context.Context, platform string) (Message, error)
		Save(ctx context.Context, message Message) error
	}
	Message struct {
		id                   string
		content              string
		recipientPhoneNumber string
		status               string
	}
)

func NewMessage(id, content, recipientPhoneNumber, status string) Message {
	return Message{
		id:                   id,
		content:              content,
		recipientPhoneNumber: recipientPhoneNumber,
		status:               status,
	}
}

func (m Message) ID() string {
	return m.id
}

func (m Message) Content() string {
	return m.content
}

func (m Message) RecipientPhoneNumber() string {
	return m.recipientPhoneNumber
}

func (m Message) Status() string {
	return m.status
}
