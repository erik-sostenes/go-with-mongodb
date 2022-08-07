package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const uri = "mongodb+srv://%s:%s@cluster0.o3tc8ee.mongodb.net/?retryWrites=true&w=majority"

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// create a new clientand connect to the server
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf(
		uri,
		os.Getenv("NoSQL_USER"),
		os.Getenv("NoSQL_PASSWORD"),
	)))

	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected and pinged")
}
