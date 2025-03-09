package app

import (
	"insider-assesment/pkg/domain/message"
	"insider-assesment/pkg/infrastructure/postgres"
	message2 "insider-assesment/pkg/infrastructure/postgres/message"
	"insider-assesment/pkg/infrastructure/redis"
	_cacheSerializer "insider-assesment/pkg/infrastructure/redis"
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
	postgresClient := postgres.NewPGPool(opt.PostgresConnectionUrl)
	redisClient := redis.NewRedisClient(opt.RedisConnectionUrl)
	_ = message2.NewService(postgresClient)
	_ = redis.NewService(redisClient, _cacheSerializer.NewSerializer())
	return &ServiceProvider{}
}
