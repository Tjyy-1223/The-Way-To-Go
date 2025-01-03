package main

import (
	"fmt"
	"time"
)

func copy3(in <-chan int) (a, b, c chan int) {
	a, b, c = make(chan int, 2), make(chan int, 2), make(chan int, 2)
	go func() {
		for {
			x := <-in
			a <- x
			b <- x
			c <- x
		}
	}()
	return
}

func fib() (out <-chan int) {
	x := make(chan int, 2)
	a, b, out := copy3(x)
	go func() {
		x <- 0
		x <- 1
		<-a
		for {
			x <- <-a + <-b
		}
	}()
	<-out
	return out
}

func mainGoFibonacci3() {
	start := time.Now()
	x := fib()
	for i := 0; i < 10; i++ {
		fmt.Println(<-x)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}
