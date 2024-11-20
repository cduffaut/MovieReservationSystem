package main

import (
	"bytes"
	"fmt"
	"net/http"
)

// POST /create-movie
// POST /new-client

// POST /new-reservation
func SendResquestPOST() {
	post_url := "http://localhost:8080/new-reservation"

	// JSON body
	body := []byte(`{
		"MovieName": "Super Man",
		"Name": "Dupont",
		"FirstName": "Jean",
		"Mail": "jeandupont@gmail.com",
		"Time": "16h00",
		"Date": "2024-11-20"
	}`)
	// Create a HTTP post request
	r, err := http.NewRequest("POST", post_url, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	r.Header.Add("Content-Type", "application/json") // inform that the content type of the request is JSON
	client := &http.Client{}
	res, err := client.Do(r) // send the client POST request
	if err != nil {
		panic(err)
	}
	fmt.Println("[+] Test: POST Request (new-reservation) Was Sended!")
	defer res.Body.Close()
}

// GET /movie-list
func SendResquestGET() {
	get_url := "http://localhost:8080/movie-list"
	// Create a HTTP GET request
	resp, err := http.Get(get_url)
	if err != nil {
		fmt.Println("Error sending GET (movie-list) request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Get (movie-list) Request failed with status code:", resp.StatusCode)
		return
	}
	fmt.Println("[+] Test: GET Request Was Sended!")
}

// DELETE /clean-outdated-movies
func SendResquestDELETE() {
	delete_url := "http://localhost:8080/clean-outdated-movies"

	// Create a HTTP DELETE request
	r, err := http.NewRequest("DELETE", delete_url, nil)
	if err != nil {
		panic(err)
	}
	r.Header.Add("Content-Type", "application/json") // inform that the content type of the request is JSON
	client := &http.Client{}
	res, err := client.Do(r) // send the client DELETE request
	if err != nil {
		panic(err)
	}
	fmt.Println("[+] Test: DELETE Request Was Sended!")
	defer res.Body.Close()
}

func main() {
	// SendResquestDELETE()
	// SendResquestGET()
	SendResquestPOST()
}
