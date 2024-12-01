package storage

import (
	"fmt"
	"log"
	"time"
)

// return false if err or if date is expired
func IsDateExpired(new_date string) (bool, time.Time, error) {
	const layout = "02-01-2006" // d-m-y

	parse_new_date, err := time.Parse(layout, new_date)
	if err != nil {
		return false, parse_new_date, err
	}

	current_date := time.Now()

	return current_date.Before(parse_new_date), parse_new_date, nil
}

func ParseMovieDate(date string) time.Time {
	var date_until time.Time
	var err error
	var res bool

	if res, date_until, err = IsDateExpired(date); err != nil {
		log.Fatal(err)
	} else if !res {
		log.Fatal("Error: Date for the new movie is expired:", date)
	}
	return date_until
}

// return true if the diffusion is before the deadline "until"
func ShowAfterDiffusionUntil(show, until time.Time) bool {
	if show.After(until) {
		fmt.Println("showtime date", show.Date, "is after the maximum date diffusion date ", until)
		return false
	}
	return true
}

// The cinema does not start to play movies from 00h to 8h30, return false if the time does not fit
func ParseTimeShow(showtime Showtime) bool {
	show_time, err := time.Parse("15:04", showtime.Time)
	if err != nil {
		fmt.Println("Invalid time format for showtime:", err, "\nSession cannot be before 08:30 or after 23:00")
		return false
	}
	// check if showtime is not too soon
	start_time, err := time.Parse("15:04", "08:30")
	if err != nil {
		log.Fatal(err)
	}
	end_time, err := time.Parse("15:04", "08:30")
	if err != nil {
		log.Fatal(err)
	}
	if !show_time.After(end_time) && !show_time.Before(start_time) {
		fmt.Println("show time", showtime, "is not in the right range: 8:30 to 23:00")
		return false
	}
	return true
}
