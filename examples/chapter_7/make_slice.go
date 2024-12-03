package main

import "fmt"

func mainMakeSlice() {
	slice1 := make([]int, 10)
	for i := 0; i < len(slice1); i++ {
		slice1[i] = i * i
	}

	// print the slice
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("%d-index's number is %d\n", i, slice1[i])
	}
	fmt.Printf("\nThe length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))
}
