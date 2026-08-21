package repository

import (
	"context"
	"fmt"

	"example.com/my_project/services/product-svc/internal/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repository struct {
	Database   *mongo.Database
	Collection *mongo.Collection
}

func NewRepository(db *mongo.Database, collection string) *Repository {
	return &Repository{
		Database:   db,
		Collection: db.Collection(collection),
	}
}

func (r *Repository) Create(
	ctx context.Context,
	product domain.Product,
) (*domain.Product, error) {
	result, err := r.Collection.InsertOne(ctx, product)
	if err != nil {
		return nil, fmt.Errorf("create product: %w", err)
	}
	id, ok := result.InsertedID.(bson.ObjectID)
	if !ok {
		return nil, fmt.Errorf("invalid product")
	}
	product.ID = id
	return &product, nil
}
