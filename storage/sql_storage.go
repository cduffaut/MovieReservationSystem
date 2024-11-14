package storage

import (
	"database/sql"

	"github.com/cduffaut/MovieReservationSystem/utils"
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

func (s *Storage) StoreMovie(movie Movie) error {
	query := `INSERT INTO MovieSession (MovieName, ClientName, ClientFirstName, ClientMail) VALUES (?, ?)`

	_, err := s.db.Exec(query, movie.MovieName, movie.ClientName, movie.ClientFirstName, movie.ClientMail)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetMovies() ([]utils.Movie, error) {
	query := `SELECT * FROM MovieSession`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movie_list []utils.Movie

	for rows.Next() {
		var movie utils.Movie
		if err := rows.Scan(&movie.MovieName, &movie.ClientName, &movie.ClientFirstName,
			&movie.ClientMail); err != nil {
			return nil, err
		}
		movie_list = append(movie_list, movie)
	}
	if err = rows.Err(); err != nil {
		return movie_list, nil
	}
	return nil, err
}
