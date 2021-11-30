package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Genre struct {
	ID    primitive.ObjectID `bson:"_id" json:"_id"`
	Genre string             `bson:"genre" json:"genre"`
}
