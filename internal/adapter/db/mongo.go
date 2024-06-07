package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	*mongo.Client
}

func NewMongoRepository() *MongoRepository {
	clientOptions := options.Client().ApplyURI("mongodb://root:root@localhost:27017") //mongodb://username:password@localhost:27017

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return &MongoRepository{
		client,
	}
}
