package main

import "fmt"

func main() {
	var age int = 16
	if age >= 18 {
		fmt.Println("You are eligible to vote")
	} else {
		fmt.Println("You are not eligible to vote")
	}

}
