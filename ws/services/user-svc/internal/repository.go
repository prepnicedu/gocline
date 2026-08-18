// /workspaces/gocline/ws/services/user-svc/internal/repository.go
package internal

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repository struct {
	Collection *mongo.Collection
	Db         *mongo.Database
}

func NewRepository(
	db *mongo.Database,
	collection string,
) *Repository {
	return &Repository{
		Db:         db,
		Collection: db.Collection(collection),
	}
}

func (r *Repository) Create(ctx context.Context, user User) (User, error) {
	result, err := r.Collection.InsertOne(ctx, user)
	if err != nil {
		return User{}, fmt.Errorf("create user: %w", err)
	}
	id, ok := result.InsertedID.(bson.ObjectID)
	if !ok {
		return User{}, fmt.Errorf("invalid id")
	}
	user.ID = id
	return user, nil
}
