package main

import (
	"fmt"
	"math"
)

type Shaper2 interface {
	Area() float32
}

type Shape struct{}

func (sh Shape) Area() float32 {
	return -1 // the shape is indetermined, so we return something impossible
}

type Square2 struct {
	side float32
	Shape
}

func (sq *Square2) Area() float32 {
	return sq.side * sq.side
}

type Rectangle2 struct {
	length, width float32
	Shape
}

func (r Rectangle2) Area() float32 {
	return r.length * r.width
}

type Circle2 struct {
	radius float32
	Shape
}

func (c *Circle2) Area() float32 {
	return math.Pi * c.radius * c.radius
}

func mainInterfaces3() {
	s := Shape{}
	r := Rectangle2{5, 3, s} // Area() of Rectangle needs a value
	q := &Square2{5, s}      // Area() of Square needs a pointer
	c := &Circle2{2.5, s}
	fmt.Println("Looping through shapes for area ...")
	// shapes := []Shaper{Shaper(r), Shaper(q), Shaper(c)}
	shapes := []Shaper{r, q, c, s}
	for n := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}
}
