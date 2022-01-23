package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Album struct {
	ID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title  string             `json:"title" bson:"title"`
	Artist string             `json:"artist" bson:"artist"`
	Price  float64            `json:"price" bson:"price"`
}
