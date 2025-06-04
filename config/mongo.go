package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {
	// Load env vars from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get Mongo URI from environment
	mongoURI := os.Getenv("MONGODB_URI")

	// Create client
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("Mongo client creation error:", err)
	}

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("MongoDB connection error:", err)
	}

	DB = client.Database("book-db")
	log.Println("✅ Connected to MongoDB!")
}

func GetCollection(name string) *mongo.Collection {
	return DB.Collection(name)
}
