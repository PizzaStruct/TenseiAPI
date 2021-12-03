package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Manga struct {
	ID          primitive.ObjectID `bson:"_id" json:"_id"`
	Title       string             `bson:"title" json:"title" validate:"required"`
	Preview     string             `bson:"preview" json:"preview"`
	Description string             `bson:"desc" json:"desc" validate:"required"`
	Genres      []string           `bson:"genres" json:"genres"`
	Year        uint32             `bson:"year" json:"year" validate:"required"`
}
