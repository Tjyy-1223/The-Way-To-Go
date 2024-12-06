package main

import (
	"fmt"
	"strconv"
)

type TwoInts2 struct {
	a, b int
}

func (tn *TwoInts2) String() string {
	return "(" + strconv.Itoa(tn.a) + "/" + strconv.Itoa(tn.b) + ")"
}

func mainMethodString() {
	two1 := &TwoInts2{12, 10}
	fmt.Printf("two1 is: %v\n", two1)
	fmt.Println("two1 is:", two1)
	fmt.Printf("two1 is: %T\n", two1)
	fmt.Printf("two1 is: %#v\n", two1)
}
