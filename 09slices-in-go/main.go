package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices")

	var slice = []int{200, 12, 55, 220, 862}
	fmt.Printf("The type of slice is: %T\n", slice)
	fmt.Println("The value of slice is:", slice)
	fmt.Println("The slice is sorted:", sort.IntsAreSorted(slice))

	sort.Ints(slice) // mutates the original slice and sorts integers
	fmt.Println("The slice is sorted:", sort.IntsAreSorted(slice))

	slice = append(slice, 12)
	fmt.Println("The sorted values of slice are:", slice)

	slice = append(slice, 76, 54)
	fmt.Println("The value of slice is:", slice)

	fmt.Println("")
	newSlice := make([]string, 2)
	newSlice[0] = "Hey"
	newSlice[1] = "How are you"
	// newSlice[2] = "Cool" // The default size is fixed to 2 so throws out of bounds error

	newSlice = append(newSlice, "Amazing", "Cool") // Extends the memory
	fmt.Println("The value of newSlice is:", newSlice)

	var slicedSlice = newSlice[1:3]                  // Gives elements between 0 to 3 indexes
	slicedSlice = append(newSlice[:2], "doing now?") // Gives elements between 0 to 2 and appends an element to it
	fmt.Println("The sliced slice from 1 to 3 is:", slicedSlice)
	fmt.Println("")

	// Removing items for slice
	var persons = []string{"Kobold", "Kimsei", "Godon", "Katie"}
	fmt.Println("Persons:", persons)
	godonIndex := 2
	persons = append(persons[:godonIndex], persons[godonIndex+1:]...)
	fmt.Println("Removed Godon from the slice:", persons)

	kimseiIndex := 1
	persons = append(persons[:kimseiIndex], persons[kimseiIndex+1:]...)
	fmt.Println("Removed Kimsei from the slice:", persons)

}
