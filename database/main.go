package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Ctx = context.TODO()

func New() *mongo.Database {
	options := options.Client()
	options.ApplyURI(os.Getenv("MONGO_URI"))
	options.SetConnectTimeout(15 * time.Second)

	client, err := mongo.Connect(Ctx, options)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(Ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return client.Database("images")
}
