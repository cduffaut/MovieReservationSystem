package main

import (
	"log"
	"net/http"
	"os"

	"github.com/cduffaut/MovieReservationSystem/postgresql"
	"github.com/cduffaut/MovieReservationSystem/request_handler"
	"github.com/cduffaut/MovieReservationSystem/storage"
	"github.com/joho/godotenv"
)

type ServerConfig struct {
	BindAddress string
}

func main() {
	godotenv.Load(".env")

	pgConfig := postgresql.Config{
		URL: os.Getenv("DATABASE_URL"),
	}

	db, err := postgresql.New(pgConfig)
	if err != nil {
		log.Fatalln("Error: Couldn't connect to Data Base:", err)
	}

	dbStorage := storage.NewSQLStorage(db)
	controller := request_handler.NewController(dbStorage)

	router := http.NewServeMux()

	router.HandleFunc("POST /create-movie", controller.CreateMovie)
	router.HandleFunc("GET /movie", controller.GetMovie)

	serverConfig := ServerConfig{
		BindAddress: os.Getenv("BIND_ADDRESS"),
	}

	httpServer := http.Server{Addr: serverConfig.BindAddress, Handler: router}

	httpServer.ListenAndServe()
}
