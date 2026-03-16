package main

import (
	"fmt"
)

type User struct {
	id    int
	name  string
	email string
}

type Person struct {
	id   int
	name string
	age  int
	city string
}

func main() {

	user := User{
		id:    1,
		name:  "Alin",
		email: "alinps@gmail.com",
	}

	user2 := User{
		id:    2,
		name:  "Aneesha",
		email: "aneesha@gmail.com",
	}

	var p Person
	p.id = 1
	p.name = "Alin"
	p.age = 24
	p.city = "Trivandrum"

	fmt.Println(user)
	fmt.Println(user.name)
	fmt.Println(user.id)
	fmt.Println(user.email)
	fmt.Println(user2.id)
	fmt.Println(user2.email)

	fmt.Println(p)

}
