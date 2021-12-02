package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/PizzaStruct/TenseiAPI/internal/pkg/models"
	"github.com/PizzaStruct/TenseiAPI/internal/pkg/repos"
	"github.com/gorilla/mux"
)

type IMangaHandler interface {
	GetManga(w http.ResponseWriter, r *http.Request)
	GetMangas(w http.ResponseWriter, r *http.Request)
	SearchMangas(w http.ResponseWriter, r *http.Request)
	InsertMangas(w http.ResponseWriter, r *http.Request)
	RemoveManga(w http.ResponseWriter, r *http.Request)
}

type MangaHandler struct {
}

func NewMangaHandler() IMangaHandler {
	return new(MangaHandler)
}

func (mh *MangaHandler) GetManga(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mangaRepo := repos.NewMangaRepo()
	manga, err := mangaRepo.GetManga(vars["id"])
	if err != nil {
		http.Error(w, "Manga not found", 404)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(manga)
}

func (mh *MangaHandler) GetMangas(w http.ResponseWriter, r *http.Request) {
	page_str := r.URL.Query().Get("page")
	page, err := strconv.Atoi(page_str)
	if err != nil {
		page = 1
	}
	mangaRepo := repos.NewMangaRepo()
	mangas := mangaRepo.GetMangas(int64(page))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(mangas)
}

func (mh *MangaHandler) SearchMangas(w http.ResponseWriter, r *http.Request) {
	page_str := r.URL.Query().Get("page")
	page, err := strconv.Atoi(page_str)
	if err != nil {
		page = 1
	}
	query := r.URL.Query().Get("query")
	mangaRepo := repos.NewMangaRepo()
	mangas := mangaRepo.SearchManga(query, int64(page))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(mangas)
}

func (mh *MangaHandler) InsertMangas(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad request", 400)
		return
	}
	var manga models.Manga
	json.Unmarshal(body, &manga)
	mangaRepo := repos.NewMangaRepo()
	if err := mangaRepo.InsertManga(&manga); err != nil {
		http.Error(w, "Bad request", 400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte("{\"status\": 201, \"message\": \"Created\"}"))
}

func (mh *MangaHandler) RemoveManga(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mangaRepo := repos.NewMangaRepo()
	mangaRepo.RemoveManga(vars["id"])
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(202)
	w.Write([]byte("{\"status\": 202, \"message\": \"Accepted\"}"))
}
