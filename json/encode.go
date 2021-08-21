package json

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type User struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Phone      int    `json:"phone"`
}

func Encrypt() {
	user := User{
		ID:         time.Now().Unix(),
		Name:       "Ade",
		Occupation: "Teacher",
		Phone:      1234567,
	}
	s, _ := json.Marshal(user)
	fmt.Printf("Encode %v : %v\n", user, string(s))

	newUser := []User{
		{
			ID:         time.Now().Unix(),
			Name:       "Ade",
			Occupation: "Programmer",
			Phone:      1234567,
		},
		{
			ID:         time.Now().Unix(),
			Name:       "Adebayo",
			Occupation: "Gopher",
			Phone:      36192765,
		},
	}

	result, err := json.Marshal(newUser)
	if err != nil {
		log.Fatal("Something Went wrong ", err)
	}
	fmt.Printf("Encode %v to %v\n", user, string(result))

}

type UserSlice struct {
	Users []User 	`json:"users"`
	Data	Data			`json:"info"`
}
type Data struct {
	Info string	`json:"info"`
}
func Encode2() {
	var users UserSlice

	users.Users = append(users.Users, []User{
		{
			ID:         time.Now().Unix(),
			Name:       "Ade",
			Occupation: "Programmer",
			Phone:      1234567,
		},
		{
			ID:         time.Now().Unix(),
			Name:       "Adebayo",
			Occupation: "Gopher",
			Phone:      36192765,
		},
		{
			ID:         time.Now().Unix(),
			Name:       "Clinton",
			Occupation: "Gopher Json",
			Phone:      36192765,
		},
	}...)

	user, err := json.MarshalIndent(users, "", "	")

	if err != nil {
		log.Fatal("Error Encoding ", err)
	}

	fmt.Printf("Encode %v\n..........\n %v\n", users, string(user))
}

func EncodeUsingMap(){
	var userMap = map[string]interface{}{
		"Id": 	time.Now().Unix(),
		"Name":	"Big Boy",
	}
	var userArray = []string{
		"good", "yes", "okay",
	}

	user, _ := json.Marshal(userMap)

	fmt.Printf("Using Map; Convert %v to %v\n", userMap, string(user))

	result, _ := json.Marshal(userArray)

	fmt.Printf("Using Map; Convert %v to %v\n", userArray, string(result))
}
func (user User) String() string {
	return fmt.Sprintf("{ ID: %v, Name: %v, Occupation: %v, Phone: %v }\n", user.ID, user.Name, user.Occupation, user.Phone)
}
