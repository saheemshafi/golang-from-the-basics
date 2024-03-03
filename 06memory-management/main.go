package main

import "fmt"

func main() {
	fmt.Println("Memory management in Golang")

	/*
		Memory management and Garbage collection in Go handles automatically after a
		threshold of a certain percentage is met that is defined in the
		`GOGC` environment variable. The default value of it is set to 100;
		Setting GOGC="off" disables the garbage collection completely
	*/

	/*
		There are two main methods for allocating memory
		new():
		  - Allocates memory but doesn't initialize it
		  - Returns the memory address
		  - Zeroed storage
		make():
		  - Allocates memory and as well as initializes it
		  - Returns the memory address
		  - Non zeroed storage
	*/
}
