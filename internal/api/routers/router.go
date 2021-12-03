package routers

import (
	"net/http"

	"github.com/PizzaStruct/TenseiAPI/internal/api/handlers"
	"github.com/gorilla/mux"
)

func InitRouter() http.Handler {
	router := mux.NewRouter().StrictSlash(true)
	mangaHandler := handlers.NewMangaHandler()

	router.HandleFunc("/mangas", mangaHandler.GetMangas).Methods("GET")
	router.HandleFunc("/mangas/{id}", mangaHandler.GetManga).Methods("GET")
	router.HandleFunc("/mangas", mangaHandler.InsertMangas).Methods("POST")
	router.HandleFunc("/mangas/{id}", mangaHandler.RemoveManga).Methods("DELETE")

	return router
}
