package main

import "fmt"

type Engine interface {
	Start()
	Stop()
}

type Car struct {
	Engine
	wheelCount int
}

func (car *Car) numberOfWheels() int {
	return car.wheelCount
}

type Mercedes struct {
	Car
}

func (m *Mercedes) sayHiToMerkel() {
	fmt.Println("Hi Angela!")
}

func (c *Car) Start() {
	fmt.Println("Car is started")
}

func (c *Car) Stop() {
	fmt.Println("Car is stopped")
}

func (c *Car) GoToWorkIn() {
	// get in car
	c.Start()
	// drive to work
	c.Stop()
	// get out of car
}

func mainInheritanceCar() {
	m := Mercedes{Car{nil, 4}}
	fmt.Println("A Mercedes has this many wheels: ", m.numberOfWheels())
	m.GoToWorkIn()
	m.sayHiToMerkel()
}
