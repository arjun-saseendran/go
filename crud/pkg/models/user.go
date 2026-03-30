package models

import (
	"context"
	"crud/pkg/config"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"username"`
	Age   int                `json:"age"`
	Place string             `json:"place"`
	Job   *Job               `json:"job"`
}

type Job struct {
	Health   string `json:"health"`
	Location string `json:"location"`
}

var collection = config.ConnectDB().Database("users").Collection("data")

func (u *User) CreateUser() *User {
	_, err := collection.InsertOne(context.TODO(), u)
	if err != nil {
		log.Fatal(err)
	}
	return u
}

func Users() []User {
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var users []User
	err = cursor.All(context.TODO(), &users)
	if err != nil {
		log.Fatal(err)
	}
	return users
}

func UserData(id string) *User {
	var user User
	cursor := collection.FindOne(context.TODO(), bson.M{"_id": id})
	cursor.Decode(&user)
	return &user
}
