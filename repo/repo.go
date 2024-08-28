package repo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const database = "tasklist"

type repo struct {
	db *mongo.Database
}

func NewRepo(db *mongo.Database) *repo {
	return &repo{db: db}
}

func NewConn(uri string) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Println("[ERROR] mongo.Connect() -->", err.Error())
		return nil, err
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Println("[ERROR] client.Ping() -->", err.Error())
		return nil, err
	}

	return client.Database(database), nil
}
