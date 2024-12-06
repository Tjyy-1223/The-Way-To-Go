package main

import "fmt"

func (p *Person) FirstName() string {
	return p.firstName
}

func (p *Person) setFirstName(newName string) {
	p.firstName = newName
}

func mainPerson2() {
	p := new(Person)
	p.setFirstName("Eric")
	fmt.Println(p.FirstName())
}
