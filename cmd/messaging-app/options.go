package main

import (
	"insider-assesment/pkg/app"
	"time"
)

func options() *app.Options {
	return &app.Options{
		PostgresConnectionUrl: "postgres://postgres:postgres@localhost:5432/insider?sslmode=disable",
		RedisConnectionUrl:    "localhost:6379",
		TriggerTime:           time.Second * 120,
		BatchSize:             2,
	}
}
