package main

import "fmt"

type employee struct {
	salary float32
}

func (this *employee) giveRaise(pct float32) {
	this.salary = this.salary * (1 + pct)
}

func mainEmployeeSalary() {
	var e = new(employee)
	e.salary = 100000

	e.giveRaise(0.4)
	fmt.Printf("Employee now makes %f rmb", e.salary)
}
