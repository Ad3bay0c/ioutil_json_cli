package json

import (
	"encoding/json"
	"fmt"
	"log"
)

func Decode() {
	var user User

	input := `{"id":1629437745,"name":"Ade","occupation":"Teacher","phone":1234567}`

	err := json.Unmarshal([]byte(input), &user)
	if err != nil {
		log.Fatal("Something went wrong ", err)
	}

	fmt.Printf("Input: %v: \n Result: %v", input, user)

	var users []User

	input = `[{"id":1629439632,"name":"Ade","occupation":"Programmer","phone":1234567},{"id":1629439632,"name":"Adebayo","occupation":"Gopher","phone":36192765}]`
	err = json.Unmarshal([]byte(input), &users)
	if err != nil {
		log.Fatal("Something went wrong ", err)
	}

	fmt.Println(users)

	input = `{"users":[{"id":1629439632,"name":"Ade","occupation":"Programmer","phone":1234567},{"id":1629439632,"name":"Adebayo","occupation":"Gopher","phone":36192765},{"id":1629439632,"name":"Clinton","occupation":"Gopher Json","phone":36192765}]}`

	var userSlice UserSlice

	err = json.Unmarshal([]byte(input), &userSlice)

	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(userSlice)

	input = `{"users":[{"id":1629439632,"name":"Ade","occupation":"Programmer","phone":1234567}], "info":{"info":"details"}}`
	err = json.Unmarshal([]byte(input), &userSlice)

	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("%+v\n",userSlice)

	var s interface{}

	input = `{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`
	err = json.Unmarshal([]byte(input), &s)

	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(s)
}