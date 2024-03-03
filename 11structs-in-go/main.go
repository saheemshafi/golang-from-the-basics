package main

import (
	"fmt"
)

func main() {
	fmt.Println("Structs in golang")

	user := User{}
	user.Id = 1
	user.Name = "John Doe"
	user.Email = "johndoe@gmail.com"

	fmt.Println("The user is:", user)
	fmt.Printf("The user in detail is: %+v\n", user)
	fmt.Println("The user's name is:", user.Name)

	anotherUser := User{2, "Jane doe", "jane@doecompany.com"}
	fmt.Println("The another user's name is:", anotherUser.Name)
}

type User struct {
	Id    int
	Name  string
	Email string
}
