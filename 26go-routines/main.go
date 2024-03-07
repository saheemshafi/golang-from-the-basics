package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Go routines")

	endpoints := []string{
		"https://google.co.in",
		"https://github.com",
		"https://facebook.com/login",
		"https://void-video-server.onrender.com/api/v1",
	}

	// for _, endpoint := range endpoints {
	// 	makeAnApiCall(endpoint) // Takes too much time. Waits for the previous call to finish
	// }

	for _, endpoint := range endpoints {
		go makeAnApiCall(endpoint) // Puts this call into a kindof separate thread managed by go runtime
		/*
			The main func has already completed execution before any of the functions started so
			nothing prints in the console
		*/
	}
}

func makeAnApiCall(endpoint string) {
	fmt.Println("Making an api call to:", endpoint)
	response, err := http.Get(endpoint)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	bodyBytes, _ := io.ReadAll(response.Body)

	fmt.Printf("The response from api: %s\n", bodyBytes)

}
