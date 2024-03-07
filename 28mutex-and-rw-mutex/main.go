package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

var items = []string{}

func main() {

	/*
		Run this program with --race flag `go run --race .`
		The problem: Multiple go routines are trying to write at same memory address relatively
		at the same time which creates the issue

		The solution: Mutex - Mutex is a locking mechanism which ensures only
		one go routine is able to access critical section (read or usually write operations)

		It prevents race conditions like 2 or more go routines trying to access/write same
		memory address at the same time
	*/

	fmt.Println("Mutex and what it solves")

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
	mutex := sync.Mutex{}

	for _, endpoint := range endpoints {
		waitGroup.Add(1)                               // adds the delta 1 to waitgroup and when it becomes 0 all go routines are released
		go makeAnApiCall(endpoint, &waitGroup, &mutex) // Puts this call into a kindof separate thread managed by go runtime
		// all api calls are launched at once - the time it takes the routine to finish may differ
	}

	waitGroup.Wait() // blocks further execution to wait for go routines to finish
	fmt.Println("The endpoint results:", items)
}

func makeAnApiCall(endpoint string, wg *sync.WaitGroup, mutex *sync.Mutex) {
	fmt.Println("Making an api call to:", endpoint)
	defer wg.Done() // decrements waitgroup counter by one so the program may not get blocked forever
	response, err := http.Get(endpoint)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	bodyBytes, _ := io.ReadAll(response.Body)

	// Solution : using mutex to lock critical actions
	mutex.Lock()
	items = append(items, string(bodyBytes))
	mutex.Unlock()

	fmt.Printf("The response from api: %s\n", bodyBytes)

}

/*
	The `sync.RWMutex` is pretty similar and it can allow read operations
	for types that are concurrent safe like maps
*/
