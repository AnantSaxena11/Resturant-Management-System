package database

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

var Client *mongo.Client

func DBinstance() *mongo.Client {
	if Client != nil {
		return Client // client already there no need to create a new instance
	}

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Erro loading the .env file")
	}

	url := os.Getenv("MONGO_URI")
	if url == "" {
		fmt.Println("MONGO_URI environment variable not set")
	}

	// context with timeout
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	// connect to mongodb

	client, err := mongo.Connect(context, options.Client().ApplyURI(url))

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context, nil)

	if err != nil {

		log.Fatal("MongoDB Pint error :", err)
	}

	fmt.Println("Connected to MongoDB")

	Client = client
	return Client
}

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("restaurant").Collection(collectionName)
	return collection
}
