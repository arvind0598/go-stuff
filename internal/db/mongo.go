package db

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	connectOnce sync.Once
	client      *mongo.Client
)

// Sets up debug logging for MongoDB commands to make life easier (we're not on prod so)
func connect() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cmdMonitor := &event.CommandMonitor{
		Started: func(_ context.Context, evt *event.CommandStartedEvent) {
			log.Print(evt.Command)
		},
	}

	clientOpts := options.Client().ApplyURI(os.Getenv("MONGO_URI")).SetMonitor(cmdMonitor)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to Database")
	}
	return client
}

// GetClient returns a singleton instance of the MongoDB client
func GetClient() *mongo.Client {
	connectOnce.Do(func() {
		client = connect()
	})
	return client
}
