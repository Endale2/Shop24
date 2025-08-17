package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func init() {
	// Load environment variables from .env if present
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, using system env")
	}
	if err := ConnectDB(); err != nil {
		log.Fatal("❌ Failed to connect to MongoDB:", err)
	}
}

func ConnectDB() error {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		return fmt.Errorf("MONGO_URI environment variable is not set")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return fmt.Errorf("Error connecting to MongoDB: %w", err)
	}

	DB = client
	log.Println("Connected to MongoDB")
	return nil
}

func GetDB() *mongo.Database {
	if DB == nil {
		log.Fatal("MongoDB client is not initialized. Call ConnectDB() first.")
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "DRPS" // Default database name
	}
	return DB.Database(dbName)
}

func GetCollection(dbName, collectionName string) *mongo.Collection {
	if DB == nil {
		log.Fatal("MongoDB client is not initialized. Call ConnectDB() first.")
	}
	return DB.Database(dbName).Collection(collectionName)
}
