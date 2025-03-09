package message

import (
	"context"
	"fmt"
)

const (
	Pending Status = iota
	Sent
	Failed
)

type (
	PostgresService interface {
		AllMessages(ctx context.Context) ([]*Message, error)
		GetMessagesByStatus(ctx context.Context, status Status, batchSize int32) ([]*Message, error)
		UpdateMessageStatus(ctx context.Context, id uint64, status Status) error
	}
	Message struct {
		id                   uint64
		content              string
		recipientPhoneNumber string
		status               Status
	}
	Status uint8
)

func NewMessage(id uint64, content, recipientPhoneNumber string, status Status) *Message {
	return &Message{
		id:                   id,
		content:              content,
		recipientPhoneNumber: recipientPhoneNumber,
		status:               status,
	}
}

func (m Message) ID() uint64 {
	return m.id
}

func (m Message) Content() string {
	return m.content
}

func (m Message) RecipientPhoneNumber() string {
	return m.recipientPhoneNumber
}

func (m Message) Status() Status {
	return m.status
}

func (s Status) String() string {
	switch s {
	case Pending:
		return "Pending"
	case Sent:
		return "Sent"
	case Failed:
		return "Failed"
	default:
		return fmt.Sprintf("Unknown Status (%d)", s)
	}
}
