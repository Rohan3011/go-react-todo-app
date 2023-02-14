package db

import (
	"context"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientInstance *mongo.Client
var mongoOnce sync.Once
var clientInstanceError error

type Collection string

const (
	TodosCollection Collection = "todos"
)

const (
	Database = "todos-api"
)

func GetMongoClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))

		client, err := mongo.Connect(context.TODO(), clientOptions)

		clientInstance = client
		clientInstanceError = err
	})

	return clientInstance, clientInstanceError
}
