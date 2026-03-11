package main

import "fmt"

func greet(name string) {
	fmt.Println("Hello ", name)
}
func add(a int, b int) int {
	var sum int = a + b
	return sum
}

func main() {
	greet("Alin")
	var result int = add(5, 4)
	fmt.Println("Result= ", result)
}
