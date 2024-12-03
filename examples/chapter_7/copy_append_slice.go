package main

import "fmt"

func mainCopyAppendSlice() {
	s1_from := []int{1, 2, 3}
	s1_to := make([]int, 10)

	n := copy(s1_to, s1_from)
	fmt.Println(s1_to)
	fmt.Printf("Copied %d elements\n", n) // n == 3

	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)
}
