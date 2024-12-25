package main

import (
	"fmt"
	"time"
)

func pump41(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 2
	}
}

func pump42(ch chan int) {
	for i := 0; ; i++ {
		ch <- i + 5
	}
}

func suck4(ch1, ch2 chan int) {
	for {
		select {
		case v := <-ch1:
			fmt.Printf("Received on channel 1: %d\n", v)
		case v := <-ch2:
			fmt.Printf("Received on channel 2: %d\n", v)
		}
	}
}

func mainGoroutineSelect() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go pump41(ch1)
	go pump42(ch2)
	go suck4(ch1, ch2)

	time.Sleep(1e9)
}
