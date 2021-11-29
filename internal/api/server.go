package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/PizzaStruct/TenseiAPI/internal/api/routers"
)

func Run() {
	go func() {
		serv := http.Server{
			Addr:    fmt.Sprintf(":%s", os.Getenv("PORT")),
			Handler: routers.InitRouter(),
		}
		if err := serv.ListenAndServe(); err != nil {
			panic(err)
		}
	}()
}
