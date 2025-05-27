package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type ResponseProduct struct {
	ID    primitive.ObjectID `json:"id"`
	Name  string             `json:"name"`
	Price float64            `json:"price"`
}
