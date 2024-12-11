package main

import "fmt"

type TopologicalGenus interface {
	Rank() int
}

func (sq *Square) Rank() int {
	return 1
}

func (r Rectangle) Rank() int {
	return 2
}

func mainAbstractInterfaces() {
	r := Rectangle{5, 3} // Area() of Rectangle needs a value
	q := &Square{5}      // Area() of Square needs a pointer

	topgen := []TopologicalGenus{r, q}
	for n, _ := range topgen {
		fmt.Println("Shape details: ", topgen[n])
		fmt.Println("Topological Genus of this shape is: ", topgen[n].Rank())
	}
}
