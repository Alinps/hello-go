package main

import "fmt"



func main() {
	x := 10
	
	p := &x
	*p = 50

	fmt.Println("x:", x)
	fmt.Println("p:", p)
	fmt.Println("*p:", *p)
}
