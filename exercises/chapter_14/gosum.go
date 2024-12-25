package main

import "fmt"

func sum(x, y int, ch chan int) {
	ch <- x + y
}

func mainGoSum() {
	c := make(chan int)
	go sum(1, 2, c)
	sum := <-c
	fmt.Println(sum)
}
