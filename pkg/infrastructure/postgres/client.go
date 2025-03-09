package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type PGPool struct {
	Conn *pgx.Conn
}

func NewPGPool(pgConnString string) *pgxpool.Pool {
	conn, err := pgxpool.New(context.Background(), pgConnString)
	if err != nil {
		log.Printf("Error connecting to the database: %v\n", err)
		return nil
	}
	log.Println("Successfully connected to the database")
	return conn
}
