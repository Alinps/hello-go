package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func main() {
	user := User{
		Name: "Alin",
		Age:  25,
	}

	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println(string(data))
	fmt.Println(data)
}
