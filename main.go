package main

import (
	"fmt"
)

func main() {
	var number []int = []int{10, 20, 30}
	var names []string = []string{"Alin", "Aneesha"}
	var num [5]int = [5]int{1, 2, 3, 4, 5}
	a := number
	a[0] = 100
	fmt.Println(number)
	fmt.Println(a)
	a = append(number, 10)
	num2 := num[1:]
	num3 := num[0:4]
	num4 := num[1:]
	fmt.Println(names)
	number = append(number, 30, 40, 50)
	fmt.Println(number)
	names = append(names, "Aleena", "Avanthika")
	fmt.Println(names)
	fmt.Println(num)
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(num4)
	fmt.Println(len(names))
	fmt.Println(cap(names))
	fmt.Println(a)
	fmt.Println(number)
}
