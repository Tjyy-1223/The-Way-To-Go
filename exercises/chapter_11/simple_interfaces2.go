package main

import "fmt"

type RSimple struct {
	i int
	j int
}

func (p *RSimple) Get() int {
	return p.j
}

func (p *RSimple) Put(u int) {
	p.j = u
}

func fI2(it Simpler) int {
	switch it.(type) {
	case *Simple:
		it.Put(5)
		return it.Get()
	case *RSimple:
		it.Put(50)
		return it.Get()
	default:
		return 99
	}
	return 0
}

func mainSimpleInterfaces2() {
	var s Simple
	fmt.Println(fI2(&s)) // &s is required because Get() is defined with a receiver type pointer
	var r RSimple
	fmt.Println(fI2(&r))
}
