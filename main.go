package main

import "fmt"

func main() {
	var num int
	var fact int = 1
	fmt.Print("Enter the number: ")
	fmt.Scan(&num)
	if num == 0 {
		fmt.Println("O has no factorial")
	} else if num == 1 {
		fmt.Println("factorial of 1 is 1")
	} else {
		for i := num; i >= 1; i-- {
			fact = fact * i
		}
		fmt.Println("Factorial of ", num, " is ", fact)
	}

}
