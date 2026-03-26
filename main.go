package main

import "fmt"

type Speaker interface {
	speak()
}
type Vehicle interface {
	start()
	stop()
	accelerate()
}

type Dog struct {
	Name string
}

type Car struct {
	FuelType string
}

type Truck struct {
	FuelType string
}

func (c Car) start() {
	fmt.Println("Car is starting...", c.FuelType)
}

func (c Car) stop() {
	fmt.Println("Car stopped...", c.FuelType)
}

func (c Car) accelerate() {
	fmt.Println("Car is accelerating", c.FuelType)
}

func (d Dog) speak() {
	fmt.Println("Barks", d.Name)
}

func (t Truck) start() {
	fmt.Println("Truck is starting...", t.FuelType)
}

func (t Truck) stop() {
	fmt.Println("Truck stopped...", t.FuelType)
}

func (t Truck) accelerate() {
	fmt.Println("Truck is accelerating", t.FuelType)
}

func main() {
	var s Speaker
	s = Dog{
		Name: "Bruno",
	}
	var c Vehicle
	var t Vehicle
	c = Car{
		FuelType: "Petrol",
	}
	t = Truck{
		FuelType: "Diesel",
	}
	c.start()
	c.stop()
	c.accelerate()
	t.start()
	t.stop()
	t.accelerate()
	s.speak()
}
