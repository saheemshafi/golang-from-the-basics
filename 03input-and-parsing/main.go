package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("IO and parsing related stuff")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("How many lines of code have you written?")
	input, err := reader.ReadString('\n')

	if err != nil {
		panic("IO: there was some issue reading the input")
	}

	fmt.Println(input)
	fmt.Printf("The type of input is %T \n", input)

	fmt.Println("Converting the string value to number for futher operations")
	linesOfCode, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
		panic("Failed to convert input to number.")
	}

	fmt.Println("Calculating number of hours coded...")
	var numberOfLinesInAnHour float64 = 35
	numberOfHours := linesOfCode / numberOfLinesInAnHour

	fmt.Println("You have coded", math.Ceil(numberOfHours), "hours")

	// strings package is used for string operations
	// math package for operation on numbers
	// strconv for conversion from string
}
