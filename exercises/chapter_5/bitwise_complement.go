package main

import "fmt"

func mainBitWise() {
	for i := 0; i < 10; i++ {
		fmt.Printf("the complement of %b is: %b \n", i, ^i)
	}
}
