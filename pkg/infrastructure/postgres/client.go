package postgres

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type PGPool struct {
	pool *pgxpool.Pool
}

func NewPGPool(ctx context.Context, connStr string) *PGPool {
	p, err := pgxpool.New(ctx, connStr)
	if err != nil {
		log.Print("Postgres connection failed!")
		panic(err)
	}
	return &PGPool{
		pool: p,
	}
}

func (p PGPool) Pool() *pgxpool.Pool {
	return p.pool
}
