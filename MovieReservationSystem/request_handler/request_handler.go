package request_handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/cduffaut/MovieReservationSystem/utils"
	"github.com/fatih/color"
	"github.com/go-playground/validator"
)

// true=content type is good, false=error
func CheckContentType(ct string, w http.ResponseWriter) bool {
	if ct != "" {
		mediaType := strings.ToLower(strings.TrimSpace(strings.Split(ct, ";")[0]))
		if mediaType != "application/json" {
			msg := "Content-Type header is not application/json"
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			return false
		}
		return true
	}
	return false
}

func ParsePostRequest(w http.ResponseWriter, r *http.Request) {
	response := new(utils.Movie)

	r.Body = http.MaxBytesReader(w, r.Body, 1048576) // Use http.MaxBytesReader to enforce a maximum read
	body, err := io.ReadAll(r.Body)
	if err != nil {
		color.Red("Error: Unable to read JSON body :")
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		color.Red("Error: Unable to marshal JSON due to")
		fmt.Println(err)
		return
	}
	// check if values of the JSON struct are good
	err = validator.New().Struct(response)
	if err != nil {
		color.Red("\nError: Validation failed due to:")
		fmt.Println(err)
		return
	}
	color.Magenta("\n[!] Printing The Post Request:")
	fmt.Println(response)
	utils.Movies = append(utils.Movies, response)

	color.Yellow("Creation de la request query mes couilles")
	query := `INSERT INTO MovieSession (MovieName, ClientName, ClientFirstName, ClientMail) VALUES (?, ?)`

	_, err = utils.DB.Exec(query, response.MovieName, response.ClientName, response.ClientFirstName, response.ClientMail)
	if err != nil {
		color.Red("ERROR WOTH EXEC")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something bad happened on the server"))
		return
	}
	color.Green("La fonction Exec OK")
	w.WriteHeader(http.StatusOK)        // writeHead sets the response status code and headers
	w.Write([]byte("[+] Movie Added!")) // write is used to send the response body
}

// Add the new received movie structure to the global tab var of struct
func HandlePostRequest(w http.ResponseWriter, r *http.Request) {
	color.Green("[+] New Post Request Received!")
	fmt.Println("Parsing the request...")
	ct := r.Header.Get("Content-Type") // Checking if the request is in the expected JSON format
	if !CheckContentType(ct, w) {
		color.Red("Error: Content-Type header is not application/json")
		return
	}
	if r.Method == "POST" {
		ParsePostRequest(w, r)
	}
}

// Return the tab movie structure with all the movies data
// func HandleGetRequest(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusOK)

// 	// Get the tab struct of MovieSession
// 	jsonData, err := json.Marshal(utils.Movies) // Encode the tab MovieSession structure
// 	if err != nil {
// 		color.Red("Error marshalling JSON:")
// 		fmt.Println(err)
// 		return
// 	}

// 	// Create an HTTP POST request with the JSON data
// 	url := os.Getenv("HTTP_BIND_ADDRESS")
// 	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
// 	if err != nil {
// 		color.Red("Error creating request:")
// 		fmt.Println(err)
// 		return
// 	}

// 	req.Header.Set("Content-Type", "application/json")

// 	// Send the HTTP request
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		color.Red("Error sending request:")
// 		fmt.Println(err)
// 		return
// 	}
// 	color.Cyan("\n[>] Content Of The Database Movie Sended\n(Response to get request):")
// 	fmt.Println(string(jsonData))
// 	defer resp.Body.Close()
// }

func scanRow(rows *sql.Rows) (*utils.Movie, error) {
	session := new(utils.Movie)
	err := rows.Scan(&session.MovieName,
		&session.ClientName,
		&session.ClientFirstName,
		&session.ClientMail,
	)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func HandleGetRequest(w http.ResponseWriter, r *http.Request) {
	query := `SELECT * FROM MovieSession`

	color.Blue("**1**:", query)
	rows, err := utils.DB.Query(query)
	color.Blue("**2**")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something bad happened on the server"))
		return
	}
	for rows.Next() {
		todo, err := scanRow(rows)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Something bad happened on the server"))
			return
		}
		utils.Movies = append(utils.Movies, todo)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utils.Movies)
}
