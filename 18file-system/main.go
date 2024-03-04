package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("File system")

	var fileName string
	fmt.Println("Write a json filename to create:")
	_, err := fmt.Scanln(&fileName)

	if err != nil {
		fmt.Println("Failed to scan input")
		panic(err)
	}

	fileName = "./" + fileName + ".json"

	file, err := os.Create(fileName)

	if err != nil {
		panic(err)
	}

	fmt.Println("Writing contents to file...")
	writeFile(file, "[{\"name\":\"Saheem Shafi\"},{\"name\":\"John Doe\"},{\"name\":\"Jane katie\"}]")
	defer file.Close()

	fmt.Println("Reading file...")
	fileContents := string(readFile(fileName))

	fmt.Println("The contents of file read are:", fileContents)
}

func writeFile(file *os.File, stringToWrite string) {
	length, err := io.WriteString(file, stringToWrite)

	if err != nil {
		panic(err)
	}

	fmt.Println("The length written to file is:", length)
}

func readFile(filePath string) []byte {
	bytes, err := os.ReadFile(filePath)

	if err != nil {
		panic(err)
	}

	return bytes
}
