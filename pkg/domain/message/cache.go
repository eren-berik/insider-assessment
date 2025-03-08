package message

import (
	"context"
	"time"
)

type (
	CacheService interface {
		Message(ctx context.Context, platform string) (Message, error)
		Save(ctx context.Context, message Message) error
	}
	Cache struct {
		id           string
		responseCode int
		sentTime     time.Time
	}
)

func NewMessageCache(id string, responseCode int, sentTime time.Time) Cache {
	return Cache{
		id:           id,
		responseCode: responseCode,
		sentTime:     sentTime,
	}
}

func (m Cache) ID() string {
	return m.id
}

func (m Cache) ResponseCode() int {
	return m.responseCode
}

func (m Cache) SentTime() time.Time {
	return m.sentTime
}
