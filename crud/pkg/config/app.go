package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB()*mongo.Client {
	err := godotenv.Load()
	if err != nil {
		log.Println("No env found!")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("URI is blank")
	}
	clientOptions :=
		options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connected successfully")
	return client
}
