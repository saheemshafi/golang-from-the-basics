package main

import (
	"fmt"
)

func main() {
	fmt.Println("Methods")

	user := User{}
	user.Id = 1
	user.Name = "John Doe"
	user.Email = "johndoe@gmail.com"

	fmt.Println("The user is:", user)
	fmt.Printf("The user in detail is: %+v\n", user)
	fmt.Println("The user's name is:", user.Name)

	anotherUser := User{2, "Jane doe", "jane@doecompany.com"}
	fmt.Println("The another user's name is:", anotherUser.Name)

	anotherUser.SendEmail()
	anotherUser.ChangeUserName("Bill Gates")
	fmt.Println("Outside change name fn:", user.Name)
	/* The name didn't changed on main struct because it was passed by value */

}

type User struct {
	Id    int
	Name  string
	Email string
}

func (user User) SendEmail() {
	fmt.Println("Sending email to:", user.Email)
}

func (user User) ChangeUserName(name string) {
	user.Name = name
	fmt.Println("Inside change name fn:", user.Name)
}
