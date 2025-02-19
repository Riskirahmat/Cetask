package db

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

var MongoClient *mongo.Client
var MongoDB *mongo.Database

var UserCollection *mongo.Collection
var ProjectCollection *mongo.Collection
var ColumnCollection *mongo.Collection
var TaskCollection *mongo.Collection
var ChecklistCollection *mongo.Collection

func ConnectMongoDB() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found, using system environment variables.")
	}

	mongoURI := os.Getenv("MONGO_URI")
	dbName := os.Getenv("MONGO_DB_NAME")

	clientOptions := options.Client().ApplyURI(mongoURI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	fmt.Println("🚀 Connected to MongoDB Atlas!")

	MongoClient = client
	MongoDB = client.Database(dbName)

	UserCollection = MongoDB.Collection("users")
	ProjectCollection = MongoDB.Collection("projects")
	ColumnCollection = MongoDB.Collection("columns")
	TaskCollection = MongoDB.Collection("tasks")
	ChecklistCollection = MongoDB.Collection("checklists")
	
}
