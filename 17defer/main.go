package main

import "fmt"

func main() {
	fmt.Println("Defer")

	/*
		The `defer` keyword in Golang defers the execution of the line to just before the returning
		from the function
	*/

	// The defered statements run in the reverse order they were defered: LIFO
	defer fmt.Println("you") //    |
	defer fmt.Println("are") //    |
	defer fmt.Println("How") // ---^

	fmt.Println("Hey")
	deferMore()

	// The control in returned the executes the other defered statements right after the return
}

func deferMore() {
	for i := 0; i < 5; i++ {
		defer fmt.Println("Defered index:", i)
		/*
			The control flow moves inside the function, defers the values
			0,1,2,3,4
		*/

	}
	// Just before the return: executed in reverse order 4,3,2,1,0
}
