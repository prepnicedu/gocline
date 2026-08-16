// ws/services/user-svc/internal/domain.go
package internal

import "go.mongodb.org/mongo-driver/v2/bson"

type User struct {
	ID   bson.ObjectID `bson:"_id,omitempty"`
	Name string        `bson:"name"`
}
