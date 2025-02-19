package models

import "go.mongodb.org/mongo-driver/bson/primitive"

func ObjectIDToHex(id primitive.ObjectID) string {
	return id.Hex()
}