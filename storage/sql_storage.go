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

type Showtime struct {
	Date string `json:"Date" validate:"required,min=10,max=10"` // Date format: yyyy-mm-dd
	Time string `json:"Time" validate:"required,min=5,max=5"`   // Time format: HH:mm
}

type Movie struct {
	MovieName      string     `json:"MovieName" validate:"required,min=1,max=30"`
	Category       string     `json:"Category" validate:"required,min=3,max=20"`
	DiffusionUntil string     `json:"DiffusionUntil" validate:"required,min=10,max=10"` // Usage: datetime=2006-01-02 (y-m-d)
	Showtimes      []Showtime `json:"Showtimes" validate:"required,dive,required"`
}

type Client struct {
	FirstName string `json:"FirstName" validate:"required,min=3,max=20"`
	Name      string `json:"Name" validate:"required,min=3,max=20"`
	Mail      string `json:"Mail" validate:"required,email"`
}

func (s *Storage) CreateTable() error {
	_, err := s.db.Exec("CREATE TABLE IF NOT EXISTS movie_list (MovieName TEXT PRIMARY KEY, Category TEXT NOT NULL,DiffusionUntil TEXT NOT NULL);")
	if err != nil {
		fmt.Println("Error during the creation of the TABLE \"movie_list\"")
		return err
	}
	_, err = s.db.Exec(`
    CREATE TABLE IF NOT EXISTS showtimes (
        ShowtimeID INTEGER PRIMARY KEY,
        MovieName TEXT NOT NULL,
        Date TEXT NOT NULL,  -- Format: yyyy-mm-dd
        Time TEXT NOT NULL,  -- Format: HH:mm
        FOREIGN KEY (MovieName) REFERENCES movie_list(MovieName) ON DELETE CASCADE
    );
`)
	if err != nil {
		fmt.Println("Error during the creation of the TABLE \"showtimes\"")
		return err
	}
	_, err = s.db.Exec("CREATE TABLE IF NOT EXISTS reservation_list(FirstName text, Name text, Mail text, Date text, Time text, MovieName text)")
	if err != nil {
		fmt.Println("Error during the creation of the TABLE \"reservation_list\"")
		return err
	}
	_, err = s.db.Exec("CREATE TABLE IF NOT EXISTS client_list(FirstName text, Name text, Mail text)")
	if err != nil {
		fmt.Println("Error during the creation of the TABLE \"client_list\"")
		return err
	}
	return nil
}

func (s *Storage) StoreMovie(movie Movie) error {
	movie_until_date := ParseMovieDate(movie.DiffusionUntil)
	query := `INSERT INTO movie_list (MovieName, Category, DiffusionUntil) VALUES ($1, $2, $3)`
	_, err := s.db.Exec(query, movie.MovieName, movie.Category, movie.DiffusionUntil)
	if err != nil {
		return err
	}
	for _, showtime := range movie.Showtimes {
		show_date := ParseMovieDate(showtime.Date)
		if !ShowAfterDiffusionUntil(show_date, movie_until_date) {
			continue
		} else if !ParseTimeShow(showtime) {
			continue
		}
		query = `INSERT INTO showtimes (MovieName, Date, Time) VALUES ($1, $2, $3)`
		_, err := s.db.Exec(query, movie.MovieName, showtime.Date, showtime.Time)
		if err != nil {
			return fmt.Errorf("failed to insert showtime: %w", err)
		}
	}
	return nil
}

func (s *Storage) StoreClient(client Client) error {
	query := `INSERT INTO client_list (FirstName, Name, Mail) VALUES ($1, $2, $3)`
	_, err := s.db.Exec(query, client.FirstName, client.Name, client.Mail)
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
	ParseMovieDate(reservation.Date)
	query := `INSERT INTO reservation_list (FirstName, Name, Mail, Date, Time, MovieName) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := s.db.Exec(query, reservation.FirstName, reservation.Name, reservation.Mail, reservation.Date, reservation.Time, reservation.MovieName)

	if err != nil {
		return err
	}
	return nil
}

// putting the new up to date movie list in the database
func (s *Storage) UpdateMovieList(up_to_date_movie_list []Movie) error {
	for _, movie := range up_to_date_movie_list {
		query := `INSERT INTO movie_list (MovieName, Category, DiffusionUntil) VALUES ($1, $2, $3))`
		_, err := s.db.Exec(query, movie.MovieName, movie.Category, movie.DiffusionUntil)
		if err != nil {
			return err
		}
	}
	return nil
}

// deleting the past movie list
func (s *Storage) DeleteMoviesList() error {
	query := `DELETE FROM movie_list`
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
		if res, _, err := IsDateExpired(movie.DiffusionUntil); err != nil || res {
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
		return movie_list, err
	}
	return movie_list, nil
}
