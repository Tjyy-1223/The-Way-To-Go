package main

import "fmt"

func mainLenCap() {
	s := make([]int, 5)
	s = s[2:4]
	fmt.Println(len(s))
	fmt.Println(cap(s))
}
