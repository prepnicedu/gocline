package domain

import "go.mongodb.org/mongo-driver/v2/bson"

type Product struct {
	ID          bson.ObjectID `bson:"_id,omitempty"`
	Title       string        `bson:"title"`
	Description string        `bson:"description"`
	ProductId   string        `bson:"product_id"`
	Qty         int32         `bson:"qty"`
}
