package storage

type StorageInterface interface {
	StoreBook(movie Movie) error
}
