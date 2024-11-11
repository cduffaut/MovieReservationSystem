package create_data_base

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // "_"=the init function of the package will be called. Go compiler will not complain if the package is not used
	"github.com/joho/godotenv"
)

var (
	DATABASE_URL, DB_DRIVER, PORT string
)

// Loading our environment variables
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error: Couldn't load the .env file")
	}
	DATABASE_URL = os.Getenv("DATABASE_URL")
	DB_DRIVER = os.Getenv("DB_DRIVER")
	PORT = os.Getenv("PORT")
}

// creating a connection with the database,
// this function returns a client using which we interact/perform operations on tables.
func DBClient() (*sql.DB, error) {
	db, err := sql.Open(DB_DRIVER, DATABASE_URL)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("[info] Connected to DB")
	return db, nil
}
