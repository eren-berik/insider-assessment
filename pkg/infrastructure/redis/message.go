package redis

import "github.com/go-redis/redis/v8"

const cacheKeyPrefixModel = "message:"

type Service struct {
	redis *redis.Client
}

func NewService(redis *redis.Client) *Service {
	return &Service{
		redis: redis,
	}
}
