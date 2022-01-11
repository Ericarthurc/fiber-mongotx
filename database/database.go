package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	Client          *mongo.Client
	ItemsCollection *mongo.Collection
	Ctx             = context.TODO()
)

func ConnectDb() {
	var err error
	Client, err = mongo.Connect(Ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal(err)
	}

	err = Client.Ping(Ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	db := Client.Database("golangy")
	ItemsCollection = db.Collection("items")
}
