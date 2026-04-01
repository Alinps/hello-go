package main

import "fmt"

func change(p *int) {
	*p = 100
}

func main() {

	var x int
	x = 1
	change(&x)

	fmt.Println(x)

}
