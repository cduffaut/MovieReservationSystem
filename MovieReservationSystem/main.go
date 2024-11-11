package main

import (
	"log"
	"net/http"
	"os"

	"github.com/cduffaut/MovieReservationSystem/create_data_base"
	"github.com/cduffaut/MovieReservationSystem/request_handler"
	"github.com/cduffaut/MovieReservationSystem/utils"
	"github.com/joho/godotenv"
)

var (
	DATABASE_URL, DB_DRIVER, PORT string
)

func LaunchApp() {
	godotenv.Load()
	bind_address := os.Getenv("BIND_ADDRESS")
	router := http.NewServeMux()

	router.HandleFunc("POST /createmovie", request_handler.HandlePostRequest)
	router.HandleFunc("GET /getmovie", request_handler.HandleGetRequest)
	utils.Server = http.Server{Addr: bind_address, Handler: router}

	utils.Server.ListenAndServe()
}

func main() {
	var err error
	// Launching DB
	utils.DB, err = create_data_base.DBClient()
	if err != nil {
		log.Fatalln("Error: Couldn't connect to Data Base:", err)
	}
	LaunchApp()
}
