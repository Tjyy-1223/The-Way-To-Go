package main

import "fmt"

func mainSimpleFunction() {
	fmt.Printf("Multiply 2 * 5 * 6 = %d\n", MultiPly3Nums(2, 5, 6))
}

func MultiPly3Nums(a int, b int, c int) int {
	return a * b * c
}
