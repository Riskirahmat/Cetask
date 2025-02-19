package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Checklist struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	TaskID primitive.ObjectID `bson:"task_id" json:"task_id"`
	Title  string             `bson:"title" json:"title"`
	IsDone bool               `bson:"is_done" json:"is_done"`
}
