package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserCollection | @desc: the user ccollection on the database
var UserCollection *mongo.Collection

// User | @desc: user model struct
type User struct {
	ID    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name  string             `json:"name" bson:"name,omitempty"`
	Email string             `json:"email" bson:"email,omitempty"`
}

// func CreateUserSchema() {
// 	// Make indexes for item collection
// 	_, _ = database.ItemsCollection.Indexes().CreateOne(database.Ctx, mongo.IndexModel{
// 		Keys: bson.D{{Key: "product", Value: bsonx.String("text")}, {Key: "serial", Value: bsonx.String("text")}, {Key: "condition", Value: bsonx.String("text")}},
// 	})
// }
