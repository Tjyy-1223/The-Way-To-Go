package main

import "fmt"

func mainCasting() {
	var n int16 = 34
	var m int32

	// compiler error: cannot use n (type int16) as type int32 in assignment
	//m = n
	m = int32(n)
	fmt.Printf("32 bit int is: %d\n", m)
	fmt.Printf("16 bit int is: %d\n", n)
}
