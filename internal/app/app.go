package app

import (
	"context"
	"fmt"
	"go-auth/internal/config"
	"go-auth/internal/db"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)



type App struct {
	Config *config.Config

	MongoClient *mongo.Client
	DB *mongo.Database
}



func New(ctx context.Context) (*App, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	mongoClient, err := db.Connect(ctx, *cfg)
	if err != nil {
		return nil, err
	}


return &App{
	Config: cfg,
	MongoClient: mongoClient.Client,
	DB: mongoClient.DB,
}, nil

}



func (app *App) Close(ctx context.Context) error {

	if app.MongoClient == nil {
		log.Println("MongoDB client is nil, skipping disconnection")
		return nil
	}

	closeCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := app.MongoClient.Disconnect(closeCtx); err != nil {
		return fmt.Errorf("failed to disconnect MongoDB client: %v", err)
	}
	return nil
}
