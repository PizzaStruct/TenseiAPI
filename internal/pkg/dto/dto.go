package dto

import "github.com/PizzaStruct/TenseiAPI/internal/pkg/models"

type RepoPageResult struct {
	TotalPages int64
	HasNext    bool
	HasPrev    bool
	Mangas     []models.Manga
}
