package storage

import "fmt"

type InMemoryStorage struct {
}

func (s *InMemoryStorage) StoreMovie(movie Movie) error {
	fmt.Println("tkt c'est stocké")
	return nil
}
