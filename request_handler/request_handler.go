package request_handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/cduffaut/MovieReservationSystem/storage"
	"github.com/go-playground/validator"
)

func NewController(storage storage.StorageInterface) *Controller {
	return &Controller{
		storage: storage,
	}
}

type Controller struct {
	storage storage.StorageInterface
}

func (c *Controller) CreateMovie(w http.ResponseWriter, r *http.Request) {
	var movie storage.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		log.Fatal(err)
	}

	if err := validator.New().Struct(movie); err != nil {
		log.Fatal(err)
	}

	if err := c.storage.StoreMovie(movie); err != nil {
		fmt.Println("Error POST (New Movie) request:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "An error occured with POST (New Movie) request",
		})
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

// add a new client to the DataBase List
func (c *Controller) NewClient(w http.ResponseWriter, r *http.Request) {
	var client storage.Client

	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		log.Fatal(err)
	}

	if err := validator.New().Struct(client); err != nil {
		log.Fatal(err)
	}

	if err := c.storage.StoreClient(client); err != nil {
		fmt.Println("Error POST (New client) request:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "An error occured with POST (New client) request",
		})
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

// Add a new client to the DataBase List
func (c *Controller) NewReservation(w http.ResponseWriter, r *http.Request) {
	var reservation storage.Reservation

	if err := json.NewDecoder(r.Body).Decode(&reservation); err != nil {
		log.Fatal(err)
	}

	if err := validator.New().Struct(reservation); err != nil {
		log.Fatal(err)
	}

	if err := c.storage.StoreReservation(reservation); err != nil {
		fmt.Println("Error POST (New Reservation) request:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "An error occured with POST (New Reservation) request",
		})
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

// Delete from the database all the outdated movies
func (c *Controller) DeleteOutdatedMovies(w http.ResponseWriter, r *http.Request) {
	if err := c.storage.CleanOutdatedMovies(); err != nil {
		fmt.Println("Error DELETE request:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "An error occured with DELETE request",
		})
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

// get back the movie struct data from the db & send a OK status if != err
func (c *Controller) GetMovie(w http.ResponseWriter, r *http.Request) {
	if movie_list, err := c.storage.GetMovies(); err != nil {
		fmt.Println("Error encoding repsonse:", err)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(movie_list)
	}
}
