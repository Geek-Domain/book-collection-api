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
	// ✅ Only load .env file locally (not on Render)
	if os.Getenv("RENDER") != "true" {
		if err := godotenv.Load(); err != nil {
			log.Println("⚠️  .env file not loaded — this is expected on Render.")
		} else {
			log.Println("✅ Loaded .env file")
		}
	}

	// ✅ Get Mongo URI from environment
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		log.Fatal("❌ MONGODB_URI is not set in environment variables")
	}

	// ✅ Create MongoDB client
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("❌ Failed to create Mongo client: %v", err)
	}

	// ✅ Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("❌ MongoDB connection error: %v", err)
	}

	// ✅ Use the database
	DB = client.Database("book-db")
	log.Println("✅ Connected to MongoDB!")
}

func GetCollection(name string) *mongo.Collection {
	return DB.Collection(name)
}
