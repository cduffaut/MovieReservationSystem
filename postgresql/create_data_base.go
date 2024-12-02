package postgresql

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // "_"=the init function of the package will be called. Go compiler will not complain if the package is not used
)

type Config struct {
	URL string
}

// Creating a connection with the database,
// This function returns a client using which we interact/perform operations on tables.
func New(cfg Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.URL)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("Info: Connected to DB!")
	return db, nil
}
