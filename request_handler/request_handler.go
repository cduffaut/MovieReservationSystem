package request_handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cduffaut/MovieReservationSystem/storage"
	"github.com/cduffaut/MovieReservationSystem/utils"
	"github.com/fatih/color"
	"github.com/go-playground/validator"
)

func NewController(storage storage.StorageInterface) *Controller {
	return &Controller{
		storage: storage,
	}
}

type Controller struct {
	storage storage.StorageInterface
}

func (c *Controller) CreateMovie(w http.ResponseWriter, r *http.Request) {
	var movie storage.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		panic(err)
	}

	if err := validator.New().Struct(movie); err != nil {
		panic(err)
	}

	fmt.Print(movie)

	if err := c.storage.StoreBook(movie); err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(map[string]string{
		"status": "OK",
	})
}

func (c *Controller) GetMovie(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}

// // Add the new received movie structure to the global tab var of struct
// func HandlePostRequest(w http.ResponseWriter, r *http.Request) {
// 	color.Green("[+] New Post Request Received!")
// 	fmt.Println("Parsing the request...")
// 	ct := r.Header.Get("Content-Type") // Checking if the request is in the expected JSON format
// 	if !CheckContentType(ct, w) {
// 		color.Red("Error: Content-Type header is not application/json")
// 		return
// 	}
// 	if r.Method == "POST" {
// 		ParsePostRequest(w, r)
// 	}
// }

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
