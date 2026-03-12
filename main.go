package main

import (
	"errors"
	"fmt"
)

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Cannot divide by zero")
	}
	return a / b, nil
}
func main() {
	var a int
	var b int
	fmt.Print("Enter the first number: ")
	fmt.Scan(&a)
	fmt.Print("Enter the second number: ")
	fmt.Scan(&b)
	result, err := divide(a, b)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Result ", result)
}
