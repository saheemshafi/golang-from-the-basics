package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Javascript Object Notation - JSON")

	users := []User{
		{
			1,
			"John Doe",
			"john@doe.com",
			"ItsASecret",
			[]User{{1, "Google", "johndoe@google.com", "AccountPassword", nil}},
		},
		{
			2,
			"Katie Donner",
			"katiedonner@google.com",
			"ItsASecretAgain",
			nil,
		},
		{
			3,
			"Bill Gates",
			"billgates@microsoft.com",
			"ShhhItsASecretAgain",
			[]User{{1, "Github", "billgates@github.com", "BillAccountPassword", nil}},
		},
		{
			4,
			"Cool Guy",
			"coolguy@metaverse.com",
			"ObviouslyItsASecretAgain",
			nil,
		},
	}

	fmt.Printf("Users:\n %+v\n", users)

	usersJson, err := EncodeJson(users)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Users json:\n%s", usersJson)
}

type User struct {
	Id       int    `json:"id"` // aliasing json
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`                  // removing fields from json output
	Accounts []User `json:"accounts,omitempty"` // omiting nil values
}

func EncodeJson(users []User) ([]byte, error) {
	// jsonOutput, err := json.Marshal(users) // json
	jsonOutput, err := json.MarshalIndent(users, "", "\t") // indented

	if err != nil {
		return nil, err
	}

	return jsonOutput, nil
}
