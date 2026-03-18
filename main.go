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
type Obj struct {
	Color string
	Shape string
}

func main() {
	jsonData := `{"Name":"Alin","Age":25,"isActive":true,"Weight":64.5}`
	jsonData2 := `{"Color":"Black","Shape":"Circle"}`
	var user User
	var obj Obj
	err := json.Unmarshal([]byte(jsonData), &user)
	err2 := json.Unmarshal([]byte(jsonData2), &obj)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	if err2 != nil {
		fmt.Println(err)
	}
	fmt.Println(user.Name)
	fmt.Println(obj.Shape)
}
