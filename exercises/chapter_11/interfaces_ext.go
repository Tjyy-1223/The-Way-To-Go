package main

import "fmt"

type Triangle struct {
	base   float32
	height float32
}

type AreaInterface interface {
	Area() float32
}

type PeriInterface interface {
	Perimeter() float32
}

func (sq *Square) Perimeter() float32 {
	return 4 * sq.side
}

func (t *Triangle) Area() float32 {
	return 0.5 * t.base * t.height
}

func mainInterfacesExt() {
	var areaIntf AreaInterface
	var periIntf PeriInterface

	sq1 := new(Square)
	sq1.side = 5
	tr1 := new(Triangle)
	tr1.base = 3
	tr1.height = 5

	areaIntf = sq1
	fmt.Printf("The square has area: %f\n", areaIntf.Area())

	periIntf = sq1
	fmt.Printf("The square has perimeter: %f\n", periIntf.Perimeter())

	areaIntf = tr1
	fmt.Printf("The triangle has area: %f\n", areaIntf.Area())
}
