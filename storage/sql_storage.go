package storage

import (
	"database/sql"
	"fmt"
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
	DiffusionUntil string `json:"DiffusionUntil" validate:"required,min=10,max=10"` // Usage: datetime=2006-01-02 (y-m-d)
}

type Client struct {
	FirstName string `json:"FirstName" validate:"required,min=3,max=20"`
	Name      string `json:"Name" validate:"required,min=3,max=20"`
	Mail      string `json:"Mail" validate:"required,email"`
}

func (s *Storage) StoreMovie(movie Movie) error {
	ParseMovie(movie.DiffusionUntil)
	query := `INSERT INTO movie_list (MovieName, Category, DiffusionUntil) VALUES (movie.MovieName, movie.Category, movie.DiffusionUntil)`

	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) StoreClient(client Client) error {
	query := `INSERT INTO client_list (FirstName, Name, Mail) VALUES (client.FirstName, client.Name, client.Mail)`

	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

type Reservation struct {
	FirstName string `json:"FirstName" validate:"required,min=3,max=20"`
	Name      string `json:"Name" validate:"required,min=3,max=20"`
	Mail      string `json:"Mail" validate:"required,email"`
	Date      string `json:"Date" validate:"required,min=10,max=10"` // Usage: datetime=2006-01-02
	Time      string `json:"Time" validate:"required,min=5,max=5"`   // Format Expected=08h00
	MovieName string `json:"MovieName" validate:"required,min=1,max=20"`
}

func (s *Storage) DoesTableExist(table_name string) bool {
	_, err := s.db.Query("SELECT * FROM " + table_name + ";")

	if err != nil {
		return false
	} else {
		return true
	}
}

func (s *Storage) StoreReservation(reservation Reservation) error {
	if res, err := IsDateExpired(reservation.Date); err != nil || res {
		panic("Error: Date for the movie reservation is not correct or up to date\nPlease refer to the waited format")
	}
	if res := s.DoesTableExist("reservation_list"); !res {
		fmt.Println("La DB Reservation n'existe pas..............")
		return nil
	}
	query := `INSERT INTO reservation_list (FirstName, Name, Mail, Date, Time, MovieName) VALUES (reservation.FirstName, reservation.Name, reservation.Mail, reservation.Date, reservation.Time, reservation.MovieName)`

	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// putting the new up to date movie list in the database
func (s *Storage) UpdateMovieList(up_to_date_movie_list []Movie) error {
	for _, movie := range up_to_date_movie_list {
		query := `INSERT INTO client_list (MovieName, Category, DiffusionUntil) VALUES (movie.MovieName, movie.Category, movie.DiffusionUntil)`

		_, err := s.db.Exec(query)
		if err != nil {
			return err
		}
		movie.MovieName = "Just To Avoid Unused Variable Error..."
	}
	return nil
}

// deleting the past movie list
func (s *Storage) DeleteMoviesList() error {
	query := `DELETE * FROM movie_list`
	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) CleanOutdatedMovies() error {
	query := `SELECT * FROM movie_list`

	rows, err := s.db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	var up_to_date_movie_list []Movie

	for rows.Next() {
		var movie Movie
		if err := rows.Scan(&movie.MovieName, &movie.Category, &movie.DiffusionUntil); err != nil {
			return err
		}
		if res, err := IsDateExpired(movie.DiffusionUntil); err != nil || res {
			up_to_date_movie_list = append(up_to_date_movie_list, movie)
		}
	}
	if err = rows.Err(); err != nil {
		return err
	} else if len(up_to_date_movie_list) == 0 {
		return nil
	}
	if err = s.DeleteMoviesList(); err != nil {
		return err
	}
	if err = s.UpdateMovieList(up_to_date_movie_list); err != nil {
		return err
	}
	return nil
}

func (s *Storage) GetMovies() ([]Movie, error) {
	if res := s.DoesTableExist("movie_list"); !res {
		fmt.Println("La DB n'existe pas..............")
		return nil, nil
	}
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
