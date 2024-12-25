package main

import (
	"fmt"
	"time"
)

func getData(ch chan string) {
	var a string
	for {
		a = <-ch
		fmt.Printf(a + " ")
	}
}

func pushData(ch chan string) {
	ch <- "hello"
	time.Sleep(3e9)
	ch <- "world"
}

func mainChannelBlock3() {
	ch := make(chan string)
	go pushData(ch)
	go getData(ch)
	time.Sleep(5e9)
}
