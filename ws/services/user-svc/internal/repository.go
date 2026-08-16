// ws/services/user-svc/internal/repository.go
package internal

import (
	"context"
	"fmt"

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

func (r *Repository) Create(ctx context.Context, user User) (User, error) {
	result, err := r.Collection.InsertOne(ctx, user)
	if err != nil {
		return User{}, fmt.Errorf("insert user: %w", err)
	}
	id, ok := result.InsertedID.(bson.ObjectID)
	if !ok {
		return User{}, fmt.Errorf(
			"unexpected inserted id type %T",
			result.InsertedID,
		)
	}
	user.ID = id
	return user, nil
}
