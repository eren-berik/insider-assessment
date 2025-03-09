package message

import (
	"context"
	"time"
)

type (
	CacheSerializer interface {
		Serialize(m Cache) ([]byte, error)
		Deserialize([]byte) (Cache, error)
	}
	CacheService interface {
		Message(ctx context.Context, id string) (Cache, error)
		Save(ctx context.Context, message Cache) error
	}
	Cache struct {
		id           string
		responseCode int
		sentTime     time.Time
	}
)

func NewMessageCache(id string, sentTime time.Time) Cache {
	return Cache{
		id:       id,
		sentTime: sentTime,
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
