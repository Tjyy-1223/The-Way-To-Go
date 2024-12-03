package main

import "fmt"

type TZ int

func mainType() {
	var a, b TZ = 3, 4
	c := a + b
	fmt.Printf("c has the value : %d", c)
}
