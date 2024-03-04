package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	fmt.Println("Consuming json")

	userJson := []byte(`
		{
			"id":1,
			"name":"John Doe",
			"email":"john@doe.com",
			"password":"ItsASecret",
			"accounts":[
				{
					"id":1,
					"name":"Google",
					"email":"johndoe@google.com",
					"password":"AccountPassword",
					"accounts":null
				}
			]
		}
	`)

	fmt.Println("User is a valid json:", json.Valid(userJson))
	DecodeJson(userJson)
}

func DecodeJson(userJson []byte) {
	// var user User
	// err := json.Unmarshal(userJson, &user) // converting json into a struct

	var kvPairs map[string]interface{}
	err := json.Unmarshal(userJson, &kvPairs) // converting json into kv pairs

	if err != nil {
		panic(err)
	}

	// fmt.Printf("%#v\n", kvPairs)
	// fmt.Printf("%+v\n", user)
	LoopOverMap(kvPairs)
}

func LoopOverMap(mapData map[string]interface{}) {
	for k, v := range mapData {
		fmt.Printf("The key is %s and the value is %v and type is %T\n", k, v, v)
	}
}

type User struct {
	Id       int    `json:"id"` // aliasing json
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`                  // removing fields from json output
	Accounts []User `json:"accounts,omitempty"` // omiting nil values
}
