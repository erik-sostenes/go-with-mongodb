package repository

import (
	"context"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	syncMongo     sync.Once
	mongoDataBase *mongo.Database
	mongoClient   *mongo.Client
	err           error
)

<<<<<<< HEAD
// NewClientMongo creates a new connection to mongoClient
// use syncOnce to create only one instance of mongoClient
func NewMongoClient(config MongoDB) (*mongo.Database, error) {
	syncMongo.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), config.ConnectTimeout*time.Second)
		defer cancel()

		mongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI(config.Dns))
		err = mongoClient.Ping(context.TODO(), readpref.Primary())
=======
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, err
	}
>>>>>>> 02-01-find-commands

		mongoDataBase = mongoClient.Database(config.DatabaseName)

		err = mongoDataBase.Client().Ping(context.TODO(), readpref.Primary())
	})
	return mongoDataBase, err
}

func NewMDB(config MongoDB) (mongoDB *mongo.Database) {
	mongoDB, err := NewMongoClient(config)
	if err != nil {
		panic(err)
	}
	return
}
