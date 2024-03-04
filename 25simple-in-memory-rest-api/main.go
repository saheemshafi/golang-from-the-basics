package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Book struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Author      *Author   `json:"author"`
	PublishedOn time.Time `json:"publishedOn"`
}

type Author struct {
	Name        string `json:"name"`
	Nationality string `json:"nationality"`
	Age         int    `json:"age"`
}

func (book *Book) IsEmpty() bool {
	if book.Name == "" || book.Author == nil {
		return true
	}
	return false
}

var books = []Book{}

func main() {
	fmt.Println("A simple in memory rest api")

	router := mux.NewRouter()

	// routes
	router.HandleFunc("/", serveHome).Methods("GET")
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books", createBook).Methods("POST")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	fmt.Println("Server is starting on port 5000...")
	log.Fatal(http.ListenAndServe(":5000", router))
}

// Controllers

func serveHome(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("content-type", "application/json")

	writer.Write([]byte(`
	{
		"message":"Server is ready to handle requests."
	}
	`))
}

func getBooks(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("content-type", "application/json")
	json.NewEncoder(writer).Encode(books)
}

func getBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("content-type", "application/json")

	params := mux.Vars(request)

	for _, book := range books {
		if book.Id == params["id"] {
			json.NewEncoder(writer).Encode(book)
			break
		}
	}

	writer.Write([]byte(`
			{
				"message":"Book was not found"
			}
		`))

}

func createBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("content-type", "application/json")

	if request.Body == nil {
		writer.Write([]byte(`
		{
			"message":"Send book details in body"
		}
	`))
		return
	}

	var book Book
	json.NewDecoder(request.Body).Decode(&book)

	if book.IsEmpty() {
		writer.Write([]byte(`
			{
				"message":"Name and author are required fields"
			}
		`))
		return
	}

	bookId, err := generateRandomId()

	if err != nil {
		writer.Write([]byte(`
			{
				"message":"Failed to generate book id"
			}
		`))
		return
	}

	book.Id = bookId
	book.PublishedOn = time.Now()
	books = append(books, book)

	json.NewEncoder(writer).Encode(book)

}

func updateBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("content-type", "application/json")

	params := mux.Vars(request)

	if request.Body == nil {
		writer.Write([]byte(`
		{
			"message":"Provide fields to update"
		}
	`))
		return
	}

	for index, book := range books {
		if book.Id == params["id"] {
			books = append(books[:index], books[index+1:]...) // Removing the book that is being edited
			var book Book
			json.NewDecoder(request.Body).Decode(&book)
			book.Id = params["id"]
			books = append(books, book) // appending the edited book
			json.NewEncoder(writer).Encode(book)
			return
		}
	}

	writer.Write([]byte(`
		{
			"message":"Book not found"
		}
	`))
}

func deleteBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("content-type", "application/json")

	params := mux.Vars(request)

	for index, book := range books {
		if book.Id == params["id"] {
			books = append(books[:index], books[index+1:]...)
			writer.Write([]byte(`
				{
					"message":"Book deleted"
				}
			`))
			return
		}
	}

	writer.Write([]byte(`
		{
			"message":"Book not found"
		}
	`))
}

func generateRandomId() (string, error) {
	timestamp := time.Now().UnixNano()

	randomBytes := make([]byte, 4)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	randomHex := hex.EncodeToString(randomBytes)
	uniqueID := fmt.Sprintf("%d%s", timestamp, randomHex)

	return uniqueID, nil
}
