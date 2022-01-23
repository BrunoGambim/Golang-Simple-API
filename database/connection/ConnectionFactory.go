package connection

import (
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	CONNECTION_STRING = "mongodb://localhost:27017/"
)

var clientInstance *mongo.Client

var clientInstanceError error

var mongoOnce sync.Once

func GetMongoClient(context context.Context) (*mongo.Client, error) {
	mongoOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(CONNECTION_STRING)
		client, err := mongo.Connect(context, clientOptions)
		if err != nil {
			clientInstanceError = err
		}

		err = client.Ping(context, nil)
		if err != nil {
			clientInstanceError = err
		}
		clientInstance = client
	})
	log.Print("Connection created")
	return clientInstance, clientInstanceError
}
