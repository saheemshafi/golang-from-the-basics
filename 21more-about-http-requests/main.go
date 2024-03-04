package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const basePath = "https://jsonplaceholder.typicode.com/posts/"

func main() {
	fmt.Println("More about http requests in go")

	// MakeGetRequest()
	// MakePostRequest()
	MakePostRequestWithFormData()
}

func MakeGetRequest() {
	response, err := http.Get(basePath + "1")

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	contentBytes, _ := io.ReadAll(response.Body)

	posts := string(contentBytes)

	fmt.Println("Post:", posts)
}
func MakePostRequest() {

	requestBody := strings.NewReader(`
	   {
		 "userId":1,
		 "title":"I just made a new post",
		 "body":"Such a long, long time ago. There existed some code which never worked"
	   }
	`)

	response, err := http.Post(basePath, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	contentBytes, _ := io.ReadAll(response.Body)

	post := string(contentBytes)
	fmt.Println("Created a post:", post)
}
func MakePostRequestWithFormData() {
	requestBody := url.Values{}

	requestBody.Add("userId", "1")
	requestBody.Add("title", "I just made a new post")
	requestBody.Add("body", "Such a long, long time ago. There existed some code which never worked")

	response, err := http.PostForm(basePath, requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	contentBytes, _ := io.ReadAll(response.Body)

	post := string(contentBytes)
	fmt.Println("Created a post using url encoded form:", post)
}
