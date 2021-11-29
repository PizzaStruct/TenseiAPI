package app

import (
	"fmt"
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

	server.Run()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	out := <-sig
	fmt.Printf("\n\rProgram stopped at %d, signal: %s\n\r", time.Now().Unix(), out.String())
}
