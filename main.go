package main

import "fmt"

func change(x *int) {
	*x = 100
}

func main() {
	a := 20
	change(&a)
	fmt.Println(a)
}
