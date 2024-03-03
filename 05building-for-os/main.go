package main

import "fmt"

func main() {
	fmt.Println("Building go to native binaries")

	/*
	 The `GOOS` environment variable is responsible for the type of
	 binary go is going to build the program into
	*/

	/*
		GOOS="windows"
		GOOS="linux"
		GOOS="darwin"
	*/

	/*
		Command exmaple : GOOS="windows" go build
		Builds the program into an .exe file
	*/
}
