package main

import "fmt"

func tel(ch chan int) {
	for i := 0; i < 15; i++ {
		ch <- i
	}
}

func mainGoroutinePanic() {
	ch := make(chan int)
	go tel(ch)
	for {
		v := <-ch
		fmt.Printf("%d ", v)
	}
}
