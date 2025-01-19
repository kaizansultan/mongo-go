package database

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func ConnectDB() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println(".env not found")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("Please set your 'MONGODB_URI' environment variable")
	}

	var err error // Jangan gunakan := untuk menghindari shadowing
	client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	// Check if connection is successful
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Could not connect to MongoDB:", err)
	}
	log.Println("Successfully connected to MongoDB")
}

func GetCollection(name string) *mongo.Collection {
	// Ambil nama database dari environment variable
	dbName := os.Getenv("MONGO_DB_NAME")
	if dbName == "" {
		log.Fatal("Please set your 'MONGO_DB_NAME' environment variable")
	}

	return client.Database(dbName).Collection(name)
}
