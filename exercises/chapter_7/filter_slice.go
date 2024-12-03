package main

import "fmt"

func mainFilterSlice() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s = Filter(s, even)
	fmt.Println(s)
}

func Filter(s []int, fn func(int) bool) []int {
	var p []int
	for _, val := range s {
		if fn(val) {
			p = append(p, val)
		}
	}
	return p
}

func even(val int) bool {
	if val%2 == 0 {
		return true
	}
	return false
}
