package main

import (
	"bytes"
	"fmt"
	"net/http"
)

type Post struct {
	MovieName       string `json:"movie_name" validate:"required"`
	ClientName      string `json:"client_name" validate:"required"`
	ClientFirstName string `json:"client_firstname" validate:"required"`
	ClientMail      string `json:"client_mail" validate:"required,email"`
}

func SendResquestGET(geturl string) {

	// Create a HTTP GET request
	resp, err := http.Get(geturl)
	if err != nil {
		fmt.Println("Error sending GET request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Request failed with status code:", resp.StatusCode)
		return
	}
	fmt.Println("[+] Test: GET Request Was Sended!")
}

func SendGoodStruct(posturl string) {
	// JSON body
	body := []byte(`{
		"movie_name": "Super Man",
		"client_name": "Dupont",
		"client_firstname": "Jean",
		"client_mail": "jeandupont@gmail.com"
	}`)
	// Create a HTTP post request
	r, err := http.NewRequest("POST", posturl, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	r.Header.Add("Content-Type", "application/json") // inform that the content type of the request is JSON
	client := &http.Client{}
	res, err := client.Do(r) // send the client POST request
	if err != nil {
		panic(err)
	}
	fmt.Println("[+] Test: GOOD POST Request Was Sended!")
	defer res.Body.Close()
}

func SendMissingElementStruct(posturl string) {
	// JSON body
	body := []byte(`{
		"movie_name": "Super Man",
		"client_name": "Dupont",
		"client_firstname": "Jean"
	}`)
	// Create a HTTP post request
	r, err := http.NewRequest("POST", posturl, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	r.Header.Add("Content-Type", "application/json") // inform that the content type of the request is JSON
	client := &http.Client{}
	res, err := client.Do(r) // send the client POST request
	if err != nil {
		panic(err)
	}
	fmt.Println("[+] Test: Missing Element POST Request Was Sended!")
	defer res.Body.Close()
}

func SendGoodStructBadFormat(posturl string) {
	// JSON body
	body := []byte(`{
		"movie_name": "Super Man",
		"client_name": 3,
		"client_firstname": "Jean",
		"client_mail": "jeandupont@gmail.com"
	}`)
	// Create a HTTP post request
	r, err := http.NewRequest("POST", posturl, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	r.Header.Add("Content-Type", "application/json") // inform that the content type of the request is JSON
	client := &http.Client{}
	res, err := client.Do(r) // send the client POST request
	if err != nil {
		panic(err)
	}
	fmt.Println("[+] Test: Bad Format POST Request Was Sended!")
	defer res.Body.Close()
}

func main() {
	// posturl := "http://localhost:5000/createmovie"
	geturl := "http://localhost:5000/getmovie"
	// SendGoodStruct(posturl)
	// SendMissingElementStruct(posturl)
	// SendGoodStructBadFormat(posturl)
	SendResquestGET(geturl)
}
