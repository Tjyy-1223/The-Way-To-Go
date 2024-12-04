package main

import (
	"exercises/chapter_9/fibo"
	"fmt"
)

var nextFibo int
var op string

func mainFibonacci() {
	op = "+"
	calls()

	fmt.Println("Change of operation from + to *")
	nextFibo = 0
	op = "*"
	calls()
}

func calls() {
	next()
	fmt.Println("...")
	next()
	fmt.Println("...")
	next()
	fmt.Println("...")
	next()
}

func next() {
	result := 0
	nextFibo++
	result = fibo.Fibonacci2(op, nextFibo)
	fmt.Printf("fibonacci(%d) is: %d\n", nextFibo, result)
}
