package app

import (
	"context"
	"insider-assesment/pkg/domain/message"
	"insider-assesment/pkg/infrastructure/postgres"
	"insider-assesment/pkg/infrastructure/redis"
	"time"
)

type (
	Options struct {
		PostgresConnectionUrl string
		RedisConnectionUrl    string
		TriggerTime           time.Duration
		BatchSize             int
	}
	ServiceProvider struct {
		postgresService message.PostgresService
		redisService    message.CacheService
	}
)

func RegisterService(opt *Options) *ServiceProvider {
	postgresClient := postgres.NewPGPool(context.Background(), opt.PostgresConnectionUrl)
	redisClient := redis.NewRedisClient(opt.RedisConnectionUrl)
	_ = postgres.NewService(postgresClient)
	_ = redis.NewService(redisClient)
	return &ServiceProvider{}
}
