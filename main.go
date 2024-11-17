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
	router.HandleFunc("POST /new-client", controller.NewClient)
	router.HandleFunc("POST /new-reservation", controller.NewReservation)
	router.HandleFunc("GET /movie-list", controller.GetMovie)
	router.HandleFunc("DELETE /clean-outdated-movies", controller.DeleteOutdatedMovies)

	serverConfig := ServerConfig{
		BindAddress: os.Getenv("BIND_ADDRESS"),
	}

	httpServer := http.Server{Addr: serverConfig.BindAddress, Handler: router}

	httpServer.ListenAndServe()
}
