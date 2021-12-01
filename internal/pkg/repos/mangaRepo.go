package repos

import (
	"context"
	"errors"
	"math"

	"github.com/PizzaStruct/TenseiAPI/internal/pkg/dto"
	"github.com/PizzaStruct/TenseiAPI/internal/pkg/models"
	"github.com/PizzaStruct/TenseiAPI/internal/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IMangaRepo interface {
	GetManga(id_hex string) (models.Manga, error)
	GetMangas(page int64) dto.RepoPageResult
	SearchManga(q string, page int64) dto.RepoPageResult
	GetMangasByGenre(genre string, page int64) dto.RepoPageResult
	InsertManga(manga *models.Manga) error
	RemoveManga(id_hex string) error
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

func (mr *MangaRepo) GetMangas(page int64) dto.RepoPageResult {
	var mangas []models.Manga
	var pageSize int64 = 100
	filter := options.Find().SetSort(bson.M{"title": 1}).SetSkip((page - 1) * pageSize).SetLimit(pageSize)
	count, _ := mr.db.Collection("mangas").CountDocuments(context.Background(), bson.M{})
	cursor, _ := mr.db.Collection("mangas").Find(context.Background(), bson.M{}, filter)
	cursor.All(context.Background(), &mangas)
	totalPages := int64(math.Ceil(float64(count) / float64(pageSize)))
	return dto.RepoPageResult{
		TotalPages: totalPages,
		HasNext:    page < totalPages,
		HasPrev:    page > 1,
		Mangas:     mangas,
	}
}

func (mr *MangaRepo) SearchManga(q string, page int64) dto.RepoPageResult {
	var mangas []models.Manga
	var pageSize int64 = 100
	opts := options.Find().SetSort(bson.M{"title": 1}).SetSkip((page - 1) * pageSize).SetLimit(pageSize)
	filter := bson.M{"title": primitive.Regex{Pattern: q, Options: "i"}}
	count, _ := mr.db.Collection("mangas").CountDocuments(context.Background(), bson.M{})
	cursor, _ := mr.db.Collection("mangas").Find(context.Background(), filter, opts)
	cursor.All(context.Background(), &mangas)
	totalPages := int64(math.Ceil(float64(count) / float64(pageSize)))
	return dto.RepoPageResult{
		TotalPages: totalPages,
		HasNext:    page < totalPages,
		HasPrev:    page > 1,
		Mangas:     mangas,
	}
}

func (mr *MangaRepo) GetMangasByGenre(genre string, page int64) dto.RepoPageResult {
	var mangas []models.Manga
	var pageSize int64 = 100
	filter := options.Find().SetSort(bson.M{"title": 1}).SetSkip((page - 1) * pageSize).SetLimit(pageSize)
	count, _ := mr.db.Collection("mangas").CountDocuments(context.Background(), bson.M{})
	cursor, _ := mr.db.Collection("mangas").Find(context.Background(), bson.M{"genres": bson.M{"$in": genre}}, filter)
	cursor.All(context.Background(), &mangas)
	totalPages := int64(math.Ceil(float64(count) / float64(pageSize)))
	return dto.RepoPageResult{
		TotalPages: totalPages,
		HasNext:    page < totalPages,
		HasPrev:    page > 1,
		Mangas:     mangas,
	}
}

func (mr *MangaRepo) InsertManga(manga *models.Manga) error {
	manga.ID = primitive.NewObjectID()
	_, err := mr.db.Collection("mangas").InsertOne(context.Background(), manga)
	return err
}

func (mr *MangaRepo) RemoveManga(id_hex string) error {
	id, err := primitive.ObjectIDFromHex(id_hex)
	if err != nil {
		return err
	}
	_, err = mr.db.Collection("mangas").DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}
