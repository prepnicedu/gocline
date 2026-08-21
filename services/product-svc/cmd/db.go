package cmd

import (
	"context"
	"fmt"

	"example.com/my_project/pkg/core"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

type DBService struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func Connect(cfg *core.Config) (*DBService, error) {
	ctx, cancel := context.WithTimeout(context.Background(), cfg.ConnectionTimeout)
	defer cancel()

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(cfg.MongoURI).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(opts)
	if err != nil {
		return nil, fmt.Errorf("error creating client: %w", err)
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf("error ping db: %w", err)

	}
	return &DBService{
		Client:   client,
		Database: client.Database(cfg.MongoDatabase),
	}, nil
}

func (s *DBService) Disconnect(ctx context.Context) error {
	if s == nil || s.Client == nil {
		return nil
	}
	return s.Client.Disconnect(ctx)
}
