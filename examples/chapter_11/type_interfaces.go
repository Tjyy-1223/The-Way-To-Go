package main

import (
	"fmt"
	"math"
)

type Square2 struct {
	side float32
}

func (sq *Square2) Area() float32 {
	return sq.side * sq.side
}

type Circle2 struct {
	radius float32
}

func (ci *Circle2) Area() float32 {
	return ci.radius * ci.radius * math.Pi
}

type Shaper2 interface {
	Area() float32
}

func mainTypeInterfaces() {
	var areaIntf Shaper2
	sq1 := new(Square2)
	sq1.side = 5

	areaIntf = sq1
	// Is Square the type of areaIntf?
	if t, ok := areaIntf.(*Square2); ok {
		fmt.Printf("The type of areaIntf is: %T\n", t)
	}
	if t, ok := areaIntf.(*Circle2); ok {
		fmt.Printf("The type of areaIntf is: %T\n", t)
	} else {
		fmt.Println("areaIntf does not contain a variable of type Circle")
	}
}
