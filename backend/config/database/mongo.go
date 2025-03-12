package database

import (
    "context"
    "fmt"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func ConnectMongoDB() {
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatal("❌ MongoDB connection failed:", err)
    }

    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal("❌ MongoDB ping failed:", err)
    }

    fmt.Println("✅ MongoDB Connected Successfully!")
    MongoClient = client
}

// GetDB returns the MongoDB database instance
func GetDB() *mongo.Database {
    return MongoClient.Database("dropshipping")
}
