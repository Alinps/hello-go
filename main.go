package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name     string
	Age      int
	isActive bool
	Weight   float64
}

func main() {
	jsonData := `{"Name":"Alin","Age":25,"isActive":true,"Weight":64.5}`
	var user User
	err := json.Unmarshal([]byte(jsonData), &user)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(user.Name)
}
