package repository

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewMongoClient(config MongoDB) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), config.ConnectTimeout *time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf(config.Uri,
		config.User,
		config.Password,
	)))
	if err != nil {
		return nil, err 
	}

	defer client.Disconnect(context.TODO())

	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, err
	}

	mongoDB := client.Database(config.DatabaseName)
	if err = mongoDB.Client().Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, err
	}

	return mongoDB, err
}

func NewMDB(config MongoDB) (mongoDB *mongo.Database) {
	mongoDB, err := NewMongoClient(config)
	if err != nil {
		panic(err)
	}
	return
}
