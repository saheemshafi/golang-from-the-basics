package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Functions in Golang")

	fmt.Println("A random number popped up:", getRandomNumber(200))

	fmt.Println("Adding two random nums:", add(int(getRandomNumber(6)), int(getRandomNumber(20))))

	integer, integerDesc := returnTwoValues()

	fmt.Println("The integer is:", integer, "and description is", integerDesc)

	func() {
		fmt.Println("IIFE: Get's immediately invoked")
	}()

	someFunction := func() string { // Anonymous function
		return "Hello World"
	}

	fmt.Println(someFunction())
}

func getRandomNumber(to int32) int32 {
	source := rand.NewSource(time.Now().UnixNano())

	generator := rand.New(source)

	return generator.Int31n(to) + 1
}

func add(x int, y int) int {
	return x + y
}

func returnTwoValues() (int, string) {
	return 12, "Returned 12"
}
