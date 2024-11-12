package storage

import "fmt"

type InMemoryStorage struct {
}

func (s *InMemoryStorage) StoreBook(movie Movie) error {
	fmt.Println("tkt c'est stocké")
	return nil
}
