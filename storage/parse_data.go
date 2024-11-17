package storage

import (
	"time"
)

// return false if err or if date is expired
func IsDateExpired(new_date string) (bool, error) {
	const layout = "2006-01-02" // y-m-d

	parse_new_date, err := time.Parse(layout, new_date)
	if err != nil {
		return false, err
	}

	current_date := time.Now()

	return current_date.Before(parse_new_date), nil
}

func ParseMovie(date string) {
	if res, err := IsDateExpired(date); err != nil {
		panic(err)
	} else if !res {
		panic("Error: Date for the new movie is expired")
	}
}
