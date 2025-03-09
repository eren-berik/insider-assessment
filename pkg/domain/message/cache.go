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
		Message(ctx context.Context, id uint64) (Cache, error)
		Save(ctx context.Context, message Cache) error
	}
	Cache struct {
		id           uint64
		responseCode int
		sentTime     time.Time
	}
)

func NewMessageCache(id uint64, responseCode int, sentTime time.Time) Cache {
	return Cache{
		id:           id,
		responseCode: responseCode,
		sentTime:     sentTime,
	}
}

func (m Cache) ID() uint64 {
	return m.id
}

func (m Cache) ResponseCode() int {
	return m.responseCode
}

func (m Cache) SentTime() time.Time {
	return m.sentTime
}
