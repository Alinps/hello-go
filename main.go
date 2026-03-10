package main

import "fmt"

func main() {
	var name string
	var age int
	var place string
	fmt.Print("Enter name, age and place: ")
	fmt.Scan(&name, &age,&place)
	fmt.Println("Name: ", name)
	fmt.Println("Age: ", age)
	fmt.Println("Place",place)
}
