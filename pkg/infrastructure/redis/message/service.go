package message

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"insider-assessment/pkg/domain/message"
)

const cacheKeyPrefixModel = "message:"

type Service struct {
	redis      *redis.Client
	serializer message.CacheSerializer
}

func NewService(redis *redis.Client, serializer message.CacheSerializer) *Service {
	return &Service{
		redis:      redis,
		serializer: serializer,
	}
}

func (r Service) Message(ctx context.Context, id string) (message.Cache, error) {
	m := message.Cache{}
	res, err := r.redis.Get(ctx, cacheKey(id)).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return m, err
		}
		return m, err
	}
	m, err = r.serializer.Deserialize([]byte(res))
	if err != nil {
		return m, err
	}
	return m, nil
}

func (r Service) Save(ctx context.Context, message message.Cache) error {
	p, err := r.serializer.Serialize(message)
	if err != nil {
		return err
	}
	if _, err = r.redis.Set(ctx, cacheKey(message.ID()), p, -1).Result(); err != nil {
		return err
	}
	return nil
}

func cacheKey(messageId string) string {
	return cacheKeyPrefixModel + messageId
}
