package app

import (
	"insider-assessment/pkg/domain/message"
	"insider-assessment/pkg/infrastructure/postgres"
	message2 "insider-assessment/pkg/infrastructure/postgres/message"
	"insider-assessment/pkg/infrastructure/redis"
	_cacheSerializer "insider-assessment/pkg/infrastructure/redis/message"
	"time"
)

type (
	Options struct {
		PostgresConnectionUrl string
		RedisConnectionUrl    string
		TriggerTime           time.Duration
		BatchSize             int32
	}
	ServiceProvider struct {
		PostgresService      message.PostgresService
		CacheService         message.CacheService
		MessageSenderService MessageSender
	}
	OptionProvider struct {
		TriggerTime time.Duration
		BatchSize   int32
	}
)

func RegisterService(opt *Options) *ServiceProvider {
	postgresClient := postgres.NewPGPool(opt.PostgresConnectionUrl)
	redisClient := redis.NewRedisClient(opt.RedisConnectionUrl)
	postgresService := message2.NewService(postgresClient)
	cacheService := _cacheSerializer.NewService(redisClient, _cacheSerializer.NewSerializer())
	messageSenderService := NewMessageSender()
	return &ServiceProvider{
		PostgresService:      postgresService,
		CacheService:         cacheService,
		MessageSenderService: *messageSenderService,
	}
}

func RegisterOptions(opt *Options) *OptionProvider {
	return &OptionProvider{
		TriggerTime: opt.TriggerTime,
		BatchSize:   opt.BatchSize,
	}
}
