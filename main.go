package main

import "fmt"

func divide(a int, b int) (int, int) {

	var quotient int = a / b
	var remainder int = a % b
	return quotient, remainder
}
func main() {
	var q int
	var r int
	q, r = divide(10, 3)
	fmt.Println("Quotien: ", q)
	fmt.Println("Remainder: ", r)
}
