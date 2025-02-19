package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
    ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Title    string             `bson:"title" json:"title"`
    Desc     string             `bson:"desc" json:"desc"`
    ColumnID primitive.ObjectID `bson:"column_id" json:"column_id"`
    UserID   primitive.ObjectID `bson:"user_id" json:"user_id"`
    Status   string             `bson:"status" json:"status"`
    Position int                `bson:"position" json:"position"`
}

