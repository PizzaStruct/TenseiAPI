package repos

import (
	"context"

	"github.com/PizzaStruct/TenseiAPI/internal/pkg/models"
	"github.com/PizzaStruct/TenseiAPI/internal/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IGenreRepo interface {
	GetGenres() []models.Genre
	InsertGenre(genre string) error
	RemoveGenre(id_hex string) error
}

type GenreRepo struct {
	db *mongo.Database
}

func NewGenreRepo() IGenreRepo {
	return &GenreRepo{mongodb.GetClient()}
}

func (gr *GenreRepo) GetGenres() []models.Genre {
	var genres []models.Genre
	opts := options.Find().SetSort(bson.M{"genre": 1})
	gr.db.Collection("genres").Find(context.Background(), bson.M{}, opts)
	return genres
}

func (gr *GenreRepo) InsertGenre(genre string) error {
	genremod := models.Genre{
		ID:    primitive.NewObjectID(),
		Genre: genre,
	}
	_, err := gr.db.Collection("genres").InsertOne(context.Background(), genremod)
	return err
}

func (gr *GenreRepo) RemoveGenre(id_hex string) error {
	id, err := primitive.ObjectIDFromHex(id_hex)
	if err != nil {
		return err
	}
	_, err = gr.db.Collection("genres").DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}
