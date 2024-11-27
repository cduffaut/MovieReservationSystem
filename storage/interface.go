package storage

type StorageInterface interface {
	StoreMovie(movie Movie) error
	GetMovies() ([]Movie, error)
	StoreClient(client Client) error
	StoreReservation(reservation Reservation) error
	CleanOutdatedMovies() error
	CreateTable() error
}
