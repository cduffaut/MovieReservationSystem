package storage

import (
	"database/sql"

	"github.com/fatih/color"
)

func NewSQLStorage(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

type Storage struct {
	db *sql.DB
}

type Movie struct {
	MovieName       string
	ClientName      string
	ClientFirstName string
	ClientMail      string
}

func (s *Storage) StoreBook(movie Movie) error {
	color.Yellow("Creation de la request query mes couilles")
	query := `INSERT INTO MovieSession (MovieName, ClientName, ClientFirstName, ClientMail) VALUES (?, ?)`

	_, err := s.db.Exec(query, movie.MovieName, movie.ClientName, movie.ClientFirstName, movie.ClientMail)
	if err != nil {
		return err
	}

	return nil
}
