package main

import (
	"fmt"
	"os"
	"time"
)

func fibonacci2(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func mainGoFibonacci2() {
	term := 25
	i := 0
	c := make(chan int)
	start := time.Now()

	go fibonacci2(term, c)

	for {
		if res, ok := <-c; ok {
			fmt.Printf("fibonacci(%d) is: %d\n", i, res)
			i++
		} else {
			end := time.Now()
			delta := end.Sub(start)
			fmt.Printf("longCalculation took this amount of time: %s\n", delta)
			os.Exit(0)
		}
	}
}
