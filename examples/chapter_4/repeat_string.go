package main

import (
	"fmt"
	"strings"
)

func mainRepeatString() {
	var origS string = "Hi there!"
	var newS string

	newS = strings.Repeat(origS, 3)
	fmt.Printf("The new repeated string is : %s \n", newS)
}
