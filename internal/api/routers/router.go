package routers

import (
	"net/http"

	"github.com/PizzaStruct/TenseiAPI/internal/api/handlers"
	"github.com/gorilla/mux"
)

func InitRouter() http.Handler {
	router := mux.NewRouter().StrictSlash(true)

	mangaRouter := router.PathPrefix("/mangas").Subrouter()
	mangaHandler := handlers.NewMangaHandler()

	mangaRouter.HandleFunc("/", mangaHandler.GetMangas).Methods("GET")
	mangaRouter.HandleFunc("/{id}", mangaHandler.GetManga).Methods("GET")
	mangaRouter.HandleFunc("/genre/{genre}", mangaHandler.GetMangasByGenre).Methods("GET")
	mangaRouter.HandleFunc("/", mangaHandler.InsertMangas).Methods("POST")
	mangaRouter.HandleFunc("/{id}", mangaHandler.RemoveManga).Methods("DELETE")

	return router
}
