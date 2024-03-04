package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const baseUrl = "https://jsonplaceholder.typicode.com/"

func main() {
	fmt.Println("Http requests in Golang")

	fmt.Println("Making an http request to:", baseUrl)
	response, err := http.Get(baseUrl + "posts")

	if err != nil {
		panic(err)
	}

	defer response.Body.Close() // Closing http requests should be done manually

	fmt.Println("Reading response body")
	posts, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(posts)
	fmt.Println("Retrieved posts from remote server:", content)

	fmt.Println("What do you want to name your name:")
	var fileName string
	fmt.Scanln(&fileName)

	fileName = "./" + fileName + ".json"
	file, err := os.Create(fileName)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	writeFile(file, content)
	fmt.Printf("Created file %v and wrote posts in that.\n", fileName)
}

func writeFile(file *os.File, stringToWrite string) {
	length, err := io.WriteString(file, stringToWrite)

	if err != nil {
		panic(err)
	}

	fmt.Println("The length written to file is:", length)
}
