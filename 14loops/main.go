package main

import "fmt"

func main() {
	fmt.Println("Loops")

	// days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println("The day is:", days[i])
	// }

	// for i := range days {
	// 	fmt.Println("The day is:", days[i])
	// }

	// for _, day := range days {
	// 	fmt.Println("The day is:", day)
	// }

	number := 0

	for number < 12 {

		if number == 10 {
			break
		}

		if number == 5 {
			number++
			continue
		}

		if number == 7 {
			goto label
		}

		fmt.Println("The number is:", number)
		number++
	}

label:
	fmt.Println("A go to label")
}
