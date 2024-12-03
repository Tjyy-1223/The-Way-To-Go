package main

import "fmt"

func mainForLoop() {
	for i := 1; i <= 15; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// 使用 goto
	i := 1
START:
	fmt.Printf("%d ", i)
	i++
	if i <= 15 {
		goto START
	}
}
