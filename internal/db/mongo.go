package db

import (
	"context"
	"fmt"
	"go-auth/internal/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type Mongo struct{ 
	Client *mongo.Client
	DB *mongo.Database
}


func Connect (ctx context.Context, config config.Config)(*Mongo, error) {

	
	connectCtx, cancel := context.WithTimeout(ctx,10* time.Second)
	defer cancel()

		clientOpts := options.Client().ApplyURI(config.MongoURI)

		// Connect to MongoDB using the provided URI and options
		client, err := mongo.Connect(ctx, clientOpts)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
		}

		// Ping the database to verify the connection is established successfully
		if err := client.Ping(connectCtx, nil); err != nil {
			return nil, fmt.Errorf("failed to ping MongoDB: %v", err)
		}

		// Access the specified database using the MongoDB client and return both the client and database instances for further operations
		database := client.Database(config.MongoDBName)

		return &Mongo{
			Client: client,
			DB: database,
		}, nil

}


func Disconnect(ctx context.Context, client *mongo.Client) error {
	// Set a timeout for the disconnection process
	connectCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

// Disconnect from MongoDB using the provided client and context
	if err := client.Disconnect(connectCtx); err != nil {
		return fmt.Errorf("failed to disconnect from MongoDB: %v", err)
	}

	return nil
}
