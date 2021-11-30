package repos

import (
	"context"
	"errors"

	"github.com/PizzaStruct/TenseiAPI/internal/pkg/models"
	"github.com/PizzaStruct/TenseiAPI/internal/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IMangaRepo interface {
	GetManga(id_hex string) (models.Manga, error)
	GetMangas(page int64) []models.Manga
}

type MangaRepo struct {
	db *mongo.Database
}

func NewMangaRepo() IMangaRepo {
	return &MangaRepo{mongodb.GetClient()}
}

func (mr *MangaRepo) GetManga(id_hex string) (models.Manga, error) {
	var manga models.Manga
	id, err := primitive.ObjectIDFromHex(id_hex)
	if err != nil {
		return manga, errors.New("manga not found")
	}
	mr.db.Collection("mangas").FindOne(context.Background(), bson.M{"_id": id}).Decode(&manga)
	return manga, nil
}

func (mr *MangaRepo) GetMangas(page int64) []models.Manga {
	var mangas []models.Manga
	var pageSize int64 = 100
	filter := options.Find().SetSort(bson.M{"_id": -1}).SetSkip((page - 1) * pageSize).SetLimit(pageSize)
	cursor, _ := mr.db.Collection("mangas").Find(context.Background(), bson.M{}, filter)
	cursor.All(context.Background(), &mangas)
	return mangas
}
