package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Search struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
}