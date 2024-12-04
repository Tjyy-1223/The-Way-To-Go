package main

import (
	"exercises/chapter_9/even"
	"fmt"
)

func mainOddEven() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("Is the integer %d even? %v\n", i, even.Even(i))
	}
}
