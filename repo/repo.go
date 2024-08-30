package repo

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

const database = "tasklist"

type repo struct {
	db *pgxpool.Pool
}

func NewRepo(db *pgxpool.Pool) *repo {
	return &repo{db: db}
}

func NewConn(uri string) (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := pgxpool.Connect(ctx, uri)
	if err != nil {
		log.Println("[ERROR] pgxpool.Connect() -->", err.Error())
		return nil, err
	}

	err = conn.Ping(ctx)
	if err != nil {
		log.Println("[ERROR] conn.Ping() -->", err.Error())
		return nil, err
	}

	return conn, nil
}
