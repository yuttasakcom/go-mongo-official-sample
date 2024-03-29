package users

import "go.mongodb.org/mongo-driver/bson/primitive"

// Users schema
type Users struct {
	ID     primitive.ObjectID `json:"_id" bson:"_id"`
	Name   string             `json:"name"`
	Age    int                `json:"age"`
	Status bool               `json:"status"`
}