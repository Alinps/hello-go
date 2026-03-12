package main

import (
	"fmt"
)

func main() {
	var numbers [3]int
	var name [4]string
	var food [3]string

	var cars [3]string = [3]string{"Mustang", "Ferrari", "lamborgini"}
	var fruits [5]string = [5]string{"apple", "orange", "lemon", "watermelon"}
	statuscode := [3]int{200, 300, 404}

	name[0] = "Alin"
	name[1] = "Aneesha"
	name[2] = "Aleena"
	name[3] = "Avanthika"

	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3

	food[0] = "mandhi"
	food[1] = "biriyani"
	food[2] = "shawarma"

	fmt.Println(numbers)
	fmt.Println(name)
	fmt.Println(food)
	fmt.Println(cars)
	fmt.Println(fruits)
	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])
	}
	for i := 0; i < len(cars); i++ {
		fmt.Printf("car %d: %s\n", i, cars[i])
	}
	for i := 0; i < len(statuscode); i++ {
		fmt.Printf("Statuscode %d: %d\n", i, statuscode[i])
	}
}
