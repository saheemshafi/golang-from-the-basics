package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Conditionals")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a number:")
	value, _ := reader.ReadString('\n')

	parsedString, err := strconv.ParseInt(strings.TrimSpace(value), 36, 64)

	if err != nil {
		fmt.Println("Error while parsing string to number:", err)
	}

	if parsedString%2 == 0 {
		fmt.Println("The number entered is even.")
	} else {
		fmt.Println("The number entered is odd.")
	}
}
