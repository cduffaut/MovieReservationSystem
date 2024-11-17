package storage

import (
	"database/sql"
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
	MovieName      string `json:"MovieName" validate:"required,min=1,max=20"`
	Category       string `json:"Category" validate:"required,min=3,max=20"`
	DiffusionUntil string `json:"DiffusionUntil" validate:"required,datetime=2006-01-02"` // Usage: datetime=2006-01-02 (y-m-d)
}

type Client struct {
	FirstName string `json:"FirstName" validate:"required,min=3,max=20"`
	Name      string `json:"Name" validate:"required,min=3,max=20"`
	Mail      string `json:"Mail" validate:"required,email"`
}

type Reservation struct {
	FirstName string `json:"FirstName" validate:"required,min=3,max=20"`
	Name      string `json:"Name" validate:"required,min=3,max=20"`
	Mail      string `json:"Mail" validate:"required,email"`
	Date      string `json:"Date" validate:"required,datetime=2006-01-02"` // Usage: datetime=2006-01-02
	Time      string `json:"Time" validate:"required,min=5,max=5"`         // Format Expected=08h00
	MovieName string `json:"MovieName" validate:"required,min=1,max=20"`
}

func (s *Storage) StoreMovie(movie Movie) error {
	ParseMovie(movie.DiffusionUntil)
	query := `INSERT INTO movie_list (MovieName, ClientName, ClientFirstName, ClientMail) VALUES (?, ?)`

	_, err := s.db.Exec(query, movie.MovieName, movie.Category, movie.DiffusionUntil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) StoreClient(client Client) error {
	query := `INSERT INTO client_list (FirstName, Name, Mail) VALUES (?, ?)`

	_, err := s.db.Exec(query, client.FirstName, client.Name, client.Mail)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) StoreReservation(reservation Reservation) error {
	query := `INSERT INTO reservation_list (FirstName, Name, Mail) VALUES (?, ?)`

	_, err := s.db.Exec(query, reservation.FirstName, reservation.Name, reservation.Mail, reservation.Date, reservation.Time, reservation.MovieName)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) GetMovies() ([]Movie, error) {
	query := `SELECT * FROM movie_list`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movie_list []Movie

	for rows.Next() {
		var movie Movie
		if err := rows.Scan(&movie.MovieName, &movie.Category, &movie.DiffusionUntil); err != nil {
			return nil, err
		}
		movie_list = append(movie_list, movie)
	}
	if err = rows.Err(); err != nil {
		return movie_list, nil
	}
	return nil, err
}
