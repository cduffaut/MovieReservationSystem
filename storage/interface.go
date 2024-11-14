package storage

import "github.com/cduffaut/MovieReservationSystem/utils"

type StorageInterface interface {
	StoreMovie(movie Movie) error
	GetMovies() ([]utils.Movie, error)
}
