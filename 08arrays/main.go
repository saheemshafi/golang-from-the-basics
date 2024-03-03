package main

import "fmt"

func main() {
	fmt.Println("Arrays in Golang")

	var array [5]string // fixed size array
	array[0] = "John Doe"
	array[1] = "Kim"
	array[2] = "Keo"
	array[3] = "Kong"
	array[4] = "Jim"

	var array2 = [2]string{"Hey", "How are you?"}

	fmt.Println("The array is:", array)
	fmt.Println("The length of this array is:", len(array)) // getting length of array

	fmt.Println("The second array is:", array2)
}
