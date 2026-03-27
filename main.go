package main

import "fmt"

type Engine struct {
	power int
}

func (e Engine) start() {
	fmt.Println("Engine started")
}

type Car struct {
	Engine
	brand string
}

func main() {
	car := Car{
		Engine: Engine{power: 500},
		brand:  "BMW",
	}

	car.start()
}
