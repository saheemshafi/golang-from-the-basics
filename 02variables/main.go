package main

import (
	"fmt"
)

// Not allowed outside of method
// tokenExpiry := 24353

// Allowed
// var tokenExpiry uint32 = 24353

const SuccessToken string = "A super secret token" // public and can't be reassigned

func main() {
	var username string = "saheemshafi"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)
	fmt.Println("")

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)
	fmt.Println("")

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)
	fmt.Println("")

	var smallFloatVal float64 = 255.32425255
	fmt.Println(smallFloatVal)
	fmt.Printf("Variable is of type: %T \n", smallFloatVal)
	fmt.Println("")

	// The default values and aliases
	var someVariable int
	fmt.Println(someVariable)
	fmt.Printf("Variable is of type: %T \n", someVariable)
	fmt.Println("")

	// implicitly declare variables
	var url = "portfolio-saheem.vercel.app"
	fmt.Println(url)
	fmt.Println("")

	// no var style to declare variables
	numberOfClicks := 1200
	fmt.Println(numberOfClicks)
	fmt.Println("")

	fmt.Println(SuccessToken)
	fmt.Printf("Variable is of type: %T \n", SuccessToken)
	fmt.Println("")

	var array = [5]int{1, 2, 3, 4, 5} // An fixed size array
	fmt.Println(array)
	fmt.Printf("Variable is of type: %T \n", array)

	fmt.Println("")
	var slice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // A variable size array also called as slice
	fmt.Println(slice)
	fmt.Printf("Variable is of type: %T \n", slice)

	fmt.Println("")
	var keyValuePairs = make(map[string]string)
	keyValuePairs["key"] = "Store key value pairs"
	fmt.Println(keyValuePairs)
	fmt.Printf("Variables is of type: %T \n", keyValuePairs)
}
