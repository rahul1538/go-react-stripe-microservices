package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

// ConnectDB connects to the MongoDB instance
func ConnectDB() error {
	// Get URI from .env
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		return fmt.Errorf("MONGO_URI is not set in .env")
	}

	// Set client options
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB (with a 10-second timeout)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	// Check the connection (Ping)
	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	DB = client
	fmt.Println("Connected to MongoDB!")
	return nil
}
