package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InitRouter() http.Handler {
	return mux.NewRouter()
}
