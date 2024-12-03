package main

import "fmt"

func mainForArray() {
	var arr [15]int
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}
	fmt.Println(arr)
}
