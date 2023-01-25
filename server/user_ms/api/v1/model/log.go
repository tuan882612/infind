package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type log struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
}