package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

func NewRedisClient(server string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: server,
	})
	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		log.Print("Redis connection failed!")
		panic(err)
	}
	return rdb
}
