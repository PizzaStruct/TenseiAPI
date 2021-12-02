package app

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	server "github.com/PizzaStruct/TenseiAPI/internal/api"
	"github.com/PizzaStruct/TenseiAPI/internal/pkg/mongodb"
	"github.com/joho/godotenv"
)

func Run() {

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	if err := mongodb.Connect(); err != nil {
		panic(err)
	}
	log.Println("Connected to mongodb")

	server.Run()
	log.Printf("Server running on port %s\n", os.Getenv("PORT"))

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	out := <-sig
	log.Printf("\n\rProgram stopped at %d, signal: %s\n\r", time.Now().Unix(), out.String())
}
