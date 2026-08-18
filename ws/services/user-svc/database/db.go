// /workspaces/gocline/ws/services/user-svc/database/db.go
package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
	"gocline.com/ws/pkg/core/config"
)

type Service struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func Connect(cfg *config.Config) (*Service, error) {
	ctx, cancel := context.WithTimeout(context.Background(), cfg.ConnectionTimeOut)
	defer cancel()

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(cfg.MongoURI).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(opts)
	if err != nil {
		return nil, fmt.Errorf("error client: %w", err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf("error client: %w", err)
	}

	return &Service{
		Client:   client,
		Database: client.Database(cfg.MongoDatabase),
	}, nil
}

func (s *Service) Disconnect(ctx context.Context) error {
	if s == nil || s.Client == nil {
		return nil
	}
	log.Println("clossing db connection..")
	return s.Client.Disconnect(ctx)
}
