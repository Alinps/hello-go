package main

import "fmt"

func main() {
	var p *string
	var v string
	v = "alin"
	p = &v
	fmt.Println("v:", v)
	fmt.Println("p:", p)
	fmt.Println("*p:", *p)
	*p = "Tony Stark"
	fmt.Println("After modifying: ", *p)
}
