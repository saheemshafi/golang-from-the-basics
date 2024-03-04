package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Modules in Golang")

	r := mux.NewRouter()

	r.HandleFunc("/", serve)

	println("Server is starting on port 4000")
	log.Fatal(http.ListenAndServe(":4000", r))
}

func serve(w http.ResponseWriter, r *http.Request) {
	// r.AddCookie(&http.Cookie{})
	// r.Method
	// r.Cookie("token")
	// r.URL.String()
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(`{
		"message":"Hey, client"
	}`))
}
