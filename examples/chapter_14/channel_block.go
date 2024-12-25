package main

import (
	"fmt"
	"time"
)

func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}

func mainChannelBlock() {
	ch1 := make(chan int)
	go pump(ch1) // pump hangs
	go suck(ch1)
	time.Sleep(3e9)
}
