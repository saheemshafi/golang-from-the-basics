package main

import (
	"fmt"
	"net/url"
)

const baseUrl = "https://jsonplaceholder.typicode.com/comments?postId=1&page=2&sort=popularity"

func main() {
	fmt.Println("URL's in Golang")

	baseUrl, _ := url.Parse(baseUrl)

	// fmt.Println("Scheme:", baseUrl.Scheme)
	// fmt.Println("Host:", baseUrl.Host)
	// fmt.Println("Path:", baseUrl.Path)
	// fmt.Println("RawQuery:", baseUrl.RawQuery)
	// fmt.Println("Query:", baseUrl.Query())

	for queryKey, query := range baseUrl.Query() {
		fmt.Println(queryKey, ":", query)
	}

	/*
		A url can be directly constructed but if we want to pass and modify it
		in other functions then a pointer should be used
	*/
	constructedUrl := &url.URL{
		Scheme:   "https",
		Host:     "github.com",
		Path:     "saheemshafi/golang-from-the-basics",
		RawQuery: "tab=issues&page=2",
	}

	fmt.Println("Constructed url:", constructedUrl.String())
}
