package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	err = Client.Ping(Ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	db := Client.Database("golangy")
	ItemsCollection = db.Collection("items")

}
