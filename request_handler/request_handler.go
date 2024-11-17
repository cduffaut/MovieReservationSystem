package request_handler

import (
	"encoding/json"
	"fmt"
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
		panic(err)
	}

	if err := validator.New().Struct(movie); err != nil {
		panic(err)
	}

	if err := c.storage.StoreMovie(movie); err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(map[string]string{
		"status": "OK",
	})
}

// add a new client to the DataBase List
func (c *Controller) NewClient(w http.ResponseWriter, r *http.Request) {
	var client storage.Client

	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		panic(err)
	}

	if err := validator.New().Struct(client); err != nil {
		panic(err)
	}

	if err := c.storage.StoreClient(client); err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(map[string]string{
		"status": "OK",
	})
}

// add a new client to the DataBase List
func (c *Controller) NewReservation(w http.ResponseWriter, r *http.Request) {
	var client storage.Client

	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		panic(err)
	}

	if err := validator.New().Struct(client); err != nil {
		panic(err)
	}

	if err := c.storage.StoreClient(client); err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(map[string]string{
		"status": "OK",
	})
}

// get back the movie struct data from the db & send a OK status if != err
func (c *Controller) GetMovie(w http.ResponseWriter, r *http.Request) {
	if movie_list, err := c.storage.GetMovies(); err != nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(movie_list)
	} else {
		fmt.Println("Error encoding repsonse:", err)
	}
}
