// ws/services/user-svc/config/db.go
package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

type DBService struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func Connect(cfg *Config) (*DBService, error) {
	ctx, cancel := context.WithTimeout(context.Background(), cfg.MongoConnectTimeout)
	defer cancel()

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(cfg.MongoURI).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(opts)
	if err != nil {
		return nil, fmt.Errorf("cannot create client: %w", err)
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf("error pint: %w", err)
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
	log.Println("closing db connection")
	return s.Client.Disconnect(ctx)
}
