package main

import "fmt"

func divide(a int, b int) (int, int) {

	var quotient int = a / b
	var remainder int = a % b
	return quotient, remainder
}
func add(a int, b int) (sum int) {
	sum = a + b
	return
}

func main() {
	var q int
	var r int
	var sumresult int
	q, r = divide(10, 3)
	sumresult = add(4, 3)
	fmt.Println("Quotien: ", q)
	fmt.Println("Remainder: ", r)
	fmt.Println("Sum: ", sumresult)
}
