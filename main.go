package main

import "fmt"

type User struct {
	name string
}

func update(u *User) {
	u.name = "updated"
}

func main() {

	var u User
	u = User{
		name: "Alin",
	}
	update(&u)
	fmt.Println(u.name)
}
