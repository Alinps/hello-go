package main

import "fmt"

type Engine struct {
	power int
}

func (e Engine) getPower() {
	fmt.Println(e.power)
}

type Car struct {
	engine Engine
	brand  string
}

func main() {
	car := Car{
		engine: Engine{power: 350},
		brand:  "toyota",
	}
	car.engine.getPower()
	fmt.Println(car.brand)
}
