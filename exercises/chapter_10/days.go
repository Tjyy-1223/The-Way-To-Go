package main

import "fmt"

type Day int

const (
	MO Day = iota
	TU
	WE
	TH
	FR
	SA
	SU
)

var dayName = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

func (day Day) String() string {
	return dayName[day]
}

func mainDay() {
	var th Day = 3
	fmt.Printf("The 3rd day is: %s\n", th)

	var day = SU
	fmt.Println(day) // prints Sunday
	fmt.Println(0, MO, 1, TU)
}
