package main

import "fmt"

func fib() func() int {
	a, b := 1, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}

func mainFibonacciClosure() {
	f := fib()
	//  Function calls are evaluated left-to-right.
	// println(f(), f(), f(), f(), f())

	for i := 0; i <= 9; i++ {
		fmt.Printf("fibonacci of %d is %d\n", i+2, f())
	}
}
