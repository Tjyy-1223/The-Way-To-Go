package main

import "fmt"

type Car struct {
	Model        string
	Manufacturer string
	BuildYear    int
}

type Cars []*Car

// Process all cars with the given function f
func (cs Cars) Process(f func(car *Car)) {
	for _, c := range cs {
		f(c)
	}
}

// FindAll Find all cars matching a given criteria.
func (cs Cars) FindAll(f func(car *Car) bool) Cars {
	cars := make([]*Car, 0)

	cs.Process(func(car *Car) {
		if f(car) {
			cars = append(cars, car)
		}
	})
	return cars
}

// Map process cars and create new data
func (cs Cars) Map(f func(car *Car) Any) []Any {
	res := make([]Any, len(cs))
	ix := 0
	cs.Process(func(car *Car) {
		res[ix] = f(car)
		ix++
	})
	return res
}

func MakeSortedAppender(manufacturers []string) (func(car *Car), map[string]Cars) {
	// Prepare maps of sorted cars
	sortedCars := make(map[string]Cars)

	for _, m := range manufacturers {
		sortedCars[m] = make([]*Car, 0)
	}
	sortedCars["Default"] = make([]*Car, 0)

	// Prepare appender function:
	// 添加到对应的制造厂商中
	appender := func(car *Car) {
		if _, ok := sortedCars[car.Manufacturer]; ok {
			sortedCars[car.Manufacturer] = append(sortedCars[car.Manufacturer], car)
		} else {
			sortedCars["Default"] = append(sortedCars["Default"], car)
		}
	}

	return appender, sortedCars
}

func mainCars() {
	// make some cars:
	ford := &Car{"Fiesta", "Ford", 2008}
	bmw := &Car{"XL 450", "BMW", 2011}
	merc := &Car{"D600", "Mercedes", 2009}
	bmw2 := &Car{"X 800", "BMW", 2008}

	// query:
	allCars := Cars([]*Car{ford, bmw, merc, bmw2})
	allNewBMWs := allCars.FindAll(func(car *Car) bool {
		return (car.Manufacturer == "BMW") && (car.BuildYear > 2010)
	})
	fmt.Println("AllCars: ", allCars)
	fmt.Println("New BMWs: ", allNewBMWs)

	manufacturers := []string{"Ford", "Aston Martin", "Land Rover", "BMW", "Jaguar"}
	sortedAppender, sortedCars := MakeSortedAppender(manufacturers)
	allCars.Process(sortedAppender)
	fmt.Println("Map sortedCars: ", sortedCars)
	BMWCount := len(sortedCars["BMW"])
	fmt.Println("We have ", BMWCount, " BMWs")
}
