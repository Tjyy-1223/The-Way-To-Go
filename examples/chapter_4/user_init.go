package main

import (
	"examples/chapter_4/trans"
	"fmt"
)

var twoPi = 4 * trans.Pi

func mainUserInit() {
	fmt.Printf("2 * Pi = %g \n", twoPi)
}
