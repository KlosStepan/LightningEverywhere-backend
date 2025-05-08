package db

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo() *mongo.Database {
	uri := os.Getenv("MONGO_URI")
	dbName := os.Getenv("DB_NAME")
	if uri == "" || dbName == "" {
		log.Fatal("Missing MONGO_URI or DB_NAME environment variables")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatalf("Mongo Connect error: %v", err)
	}

	// Ping
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Mongo Ping error: %v", err)
	}

	log.Println("✅ Connected to MongoDB")
	return client.Database(dbName)
}
