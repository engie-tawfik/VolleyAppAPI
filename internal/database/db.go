package database

import (
	"context"
	"os"
	"volleyapp/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection
var Client *mongo.Client

func Connect() {
	mongoUri := os.Getenv("MONGO_URI")
	if mongoUri == "" {
		config.Logger.Error("Database uri not found")
	}
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoUri).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	Client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		config.Logger.Error("No database connection. Error: " + err.Error())
	}
	// Send a ping to confirm a successful connection
	if err := Client.Database("volleyapp").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		config.Logger.Error("No database connection. Error: " + err.Error())
	}
	Collection = Client.Database("volleyapp").Collection("teams")
	config.Logger.Info("Successfully connected to DB")
}

func Disconnect(c context.Context) {
	Client.Disconnect(c)
}
