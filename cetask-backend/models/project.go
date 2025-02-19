package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Project struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name   string             `bson:"name" json:"name"`
	UserID primitive.ObjectID `bson:"user_id" json:"user_id"`
}
