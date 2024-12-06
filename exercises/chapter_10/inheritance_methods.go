package main

import "fmt"

type Base struct {
	id string
}

func (b *Base) Id() string {
	return b.id
}

func (b *Base) SetId(id string) {
	b.id = id
}

type Person2 struct {
	Base
	FirstName string
	LastName  string
}

type Employee struct {
	Person2
	salary float32
}

func mainInheritanceMethods() {
	idjb := Base{"007"}
	jb := Person2{idjb, "James", "Bond"}
	e := Employee{jb, 10000.}

	fmt.Printf("Id of our hero: %v\n", e.Id())
	e.SetId("007B")
	fmt.Printf("Id of our hero: %v\n", e.Id())
}
