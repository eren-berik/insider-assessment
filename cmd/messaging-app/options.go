package main

import (
	"insider-assesment/pkg/app"
	"time"
)

func options() *app.Options {
	return &app.Options{
		PostgresConnectionUrl: "postgres://postgres:postgres@postgres:5432/insider?sslmode=disable",
		RedisConnectionUrl:    "redis:6379",
		TriggerTime:           time.Second * 120,
		BatchSize:             2,
	}
}
