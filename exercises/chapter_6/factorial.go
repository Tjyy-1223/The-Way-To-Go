package main

import "fmt"

func mainFactorial() {
	for i := uint64(0); i < uint64(30); i++ {
		fmt.Printf("Factorial of %d is %d\n", i, Factorial(i))
	}
}

func Factorial(n uint64) (fac uint64) {
	fac = 1
	if n > 0 {
		fac = Factorial(n-1) * n
		return
	}
	return
}
