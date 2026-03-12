package main

import (
	"fmt"
	"reflect"
)

func main() {
	var name string = "Alin"
	var age int = 24
	var isActive bool = true
	var a float64 = 25.9
	var b int = 25
	fmt.Println("type of name:", reflect.TypeOf(name))
	fmt.Println("type of age: ", reflect.TypeOf(age))
	fmt.Println("type of isActive: ", reflect.TypeOf(isActive))
	fmt.Println("sum of float and int after explicit type conversion: ", float64(b)+a)
}
