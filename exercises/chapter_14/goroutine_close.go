package main

import "fmt"

func tel2(ch chan int) {
	for i := 0; i < 15; i++ {
		ch <- i
	}
	close(ch) // if this is ommitted: panic: all goroutines are asleep - deadlock!
}

func mainGoroutineClose() {
	ch := make(chan int)
	go tel2(ch)
	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Printf("%d ", v)
	}
}
