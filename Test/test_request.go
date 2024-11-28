package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

// POST /create-movie
func POSTCreateMovie() {
	post_url := "http://localhost:8080/create-movie"

	// JSON body
	body := []byte(`{
		"MovieName": "Inception",
		"Category": "Science Fiction",
		"DiffusionUntil": "31-12-2024",
		"Showtimes": [
		  {"Date": "28-11-2024", "Time": "14:00"},
		  {"Date": "28-11-2024", "Time": "17:30"},
		  {"Date": "28-11-2024", "Time": "20:00"}
		]
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
	fmt.Println("[+] Test: POST Request (create-movie) Was Sended!")
	defer res.Body.Close()
}

// POST /new-client

// POST /new-reservation
func SendResquestPOST() {
	post_url := "http://localhost:8080/new-reservation"

	// JSON body
	body := []byte(`{
		"MovieName": "La vache",
		"Name": "Laura",
		"FirstName": "Schefer",
		"Mail": "laulau@gmail.com",
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
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error during the reading of the response:", err)
		return
	}

	// Afficher la réponse brute
	fmt.Println("Brut response from GET:", string(body))
	fmt.Println("[+] Test: GET Request Was Sended!")
}

// GET /movie-list
func GET_() {
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
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error during the reqding of the response:", err)
		return
	}

	// Display brut response
	fmt.Println("Brut response from GET:", string(body))
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
	// SendResquestPOST()
	POSTCreateMovie()
}
