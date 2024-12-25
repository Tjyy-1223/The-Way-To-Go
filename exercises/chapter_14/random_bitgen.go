package main

import "fmt"

func mainRandomBitGen() {
	c := make(chan int)

	// consumer:
	go func() {
		for {
			fmt.Print(<-c, " ")
		}
	}()

	// producer:
	for {
		select { // random select
		case c <- 0:
		case c <- 1:
		}
	}
}
