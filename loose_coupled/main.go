package main

import "fmt"

type Engine interface {
	start()
}

type PetrolEngine struct{}

func (p PetrolEngine) start() {
	fmt.Println("Petrol engine started")
}

type DieselEngine struct{}

func (d DieselEngine) start() {
	fmt.Println("Diesel engine started")
}

type Car struct {
	engine Engine
}

func (c Car) startcar() {
	c.engine.start()
}

func main() {
	var petrol Engine
	var diesel Engine

	petrol = PetrolEngine{}
	diesel = DieselEngine{}

	var car1 Car

	car1 = Car{
		engine: petrol,
	}
	car1.startcar()

	var car2 Car

	car2 = Car{
		engine: diesel,
	}

	car2.startcar()
}
