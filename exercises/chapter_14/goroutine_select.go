package main

import "fmt"

func tel3(ch, newCh chan int) {
	for i := 0; i < 15; i++ {
		ch <- i
	}
	newCh <- 0
}

func mainGoroutineSelect() {
	ch := make(chan int)
	newCh := make(chan int)

	go tel3(ch, newCh)
	ok := false
	for {
		select {
		case <-newCh:
			ok = true
			break
		case v := <-ch:
			fmt.Printf("%d ", v)
		}
		if ok {
			break
		}
	}
}
