package config

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoCon struct {
	Ctx        context.Context
	Client     *mongo.Client
	Collection *mongo.Collection
	Cancel     context.CancelFunc
}

func WithDbCon(callback func(*MongoCon)) error {
	url := os.Getenv("MONGO_URL")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	defer client.Disconnect(ctx)
	if err != nil {
		return err
	}

	err = client.Connect(ctx)
	if err != nil {
		return err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}
	collection := client.Database("arxiv-insanity").Collection("project")
	callback(&MongoCon{ctx, client, collection, cancel})
	return nil
}
