package main

import "fmt"

type Speaker interface {
	speak()
}

type Dog struct {
	Name string
}

func (d Dog) speak() {
	fmt.Println("Barks", d.Name)
}

func main() {
	var s Speaker
	s = Dog{
		Name: "Bruno",
	}
	s.speak()
}
