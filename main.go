package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	isActive bool    `json:"isactive"`
	Weight   float64 `json:"weight"`
}

func main() {
	user := User{
		Name:     "Alin",
		Age:      25,
		isActive: true,
		Weight:   64.75,
	}

	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println(string(data))
	fmt.Println(data)
}
