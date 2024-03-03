package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch case")

	fmt.Println("Rolling a dice...")
	source := rand.NewSource(time.Now().UnixNano())

	generator := rand.New(source)

	diceResult := generator.Intn(6) + 1

	fmt.Println("The dice was rolled and the result is:", diceResult)

	switch diceResult {
	case 1:
		fmt.Println("You got 1 and can open")
	case 2:
		fmt.Println("You can move 2 places")
	case 3:
		fmt.Println("You can move 3 places")
	case 4:
		fmt.Println("You can move 4 places")
	case 5:
		fmt.Println("You can move 5 places")
		// fallthrough // can optionally specific fallthrough if want to go through all remaining cases
	case 6:
		fmt.Println("You can move 6 places and roll again")
	default:
		fmt.Println("WTH happened!")
	}
}
