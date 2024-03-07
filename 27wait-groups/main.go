package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

func main() {

	fmt.Println("Go routines")

	endpoints := []string{
		"https://jsonplaceholder.typicode.com/users/1",
		"https://jsonplaceholder.typicode.com/todos/1",
		"https://void-video-server.onrender.com/api/v1",
	}

	/*
		A wait group is used to wait for all the go routines currently launched
		finish their work
	*/
	waitGroup := sync.WaitGroup{}
	defer waitGroup.Wait() // waits for the delta of wait group to become 0

	for _, endpoint := range endpoints {
		waitGroup.Add(1)                       // adds the delta 1 to waitgroup and when it becomes 0 all go routines are released
		go makeAnApiCall(endpoint, &waitGroup) // Puts this call into a kindof separate thread managed by go runtime
		// all api calls are launched at once - the time it takes the routine to finish may differ
	}
}

func makeAnApiCall(endpoint string, wg *sync.WaitGroup) {
	fmt.Println("Making an api call to:", endpoint)
	defer wg.Done() // decrements waitgroup counter by one so the program may not get blocked forever
	response, err := http.Get(endpoint)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	bodyBytes, _ := io.ReadAll(response.Body)

	fmt.Printf("The response from api: %s\n", bodyBytes)
}
