package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Column struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ProjectID primitive.ObjectID `bson:"project_id" json:"project_id"`
	Name      string             `bson:"name" json:"name"`
	UserID    primitive.ObjectID `bson:"user_id" json:"user_id"`
	Position  int                `bson:"position" json:"position"`
}
