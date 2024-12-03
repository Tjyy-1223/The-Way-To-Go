package main

import "fmt"

var s []int

func mainMagnifySlice() {
	s = []int{1, 2, 3}
	fmt.Println("The length of s before enlarging is:", len(s))
	fmt.Println(s)

	s = enlarge(s, 5)
	fmt.Println("The length of s after enlarging is:", len(s))
	fmt.Println(s)
}

func enlarge(s []int, factor int) []int {
	ns := make([]int, len(s)*factor)
	// fmt.Println("The length of ns  is:", len(ns))
	copy(ns, s)
	s = ns
	return s
}
