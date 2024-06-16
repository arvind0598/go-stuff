package db

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	connectOnce sync.Once
	client      *mongo.Client
)

func connect() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to Database")
	}
	return client
}

func GetClient() *mongo.Client {
	connectOnce.Do(func() {
		client = connect()
	})
	return client
}
