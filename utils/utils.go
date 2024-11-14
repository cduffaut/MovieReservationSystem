package utils

type Movie struct {
	MovieName       string `json:"movie_name" validate:"required"`
	ClientName      string `json:"client_name" validate:"required"`
	ClientFirstName string `json:"client_firstname" validate:"required"`
	ClientMail      string `json:"client_mail" validate:"required,email"`
}
