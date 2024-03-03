package main

func main() {
	println("Pointers in golang")

	/*
		A pointer is simply a memory address of a value
		We can reference a memory address by prefixing the variable/value by `&`
		We can deference the value from the reference by `*`
	*/

	// Example
	personName := "John doe"

	var personRef *string = &personName

	println("The person is present on", personRef, "in memory")
	println("The value the personRef references", *personRef, "at", personRef, "memory address")

	/*
		A pointer makes sure that a ref of the memory is passed down the line
		so that the actual value is modified/mutated/operated-on.

		Most of the variables in Go are passed by value i.e the copies of actual
		variables are passed which don't effect the actual variable
	*/

	changeName(personName)
	println("The name didn't changes as it was passed as a copy:", personName)

	changeNameByPointer(&personName)
	println("The actual name changed as a ref was passed:", personName)
}

func changeName(name string) {
	name = "Jane doe"
}

func changeNameByPointer(personName *string) {
	*personName = "Jane Doe"
}
