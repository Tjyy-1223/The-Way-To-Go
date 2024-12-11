package main

import "fmt"

func mainFloat64Main() {
	f1 := newFloat64Array()
	f1.Fill(10)
	fmt.Printf("Before sorting %s\n", f1)

	Sort(f1)
	fmt.Printf("After sorting %s\n", f1)

	if IsSorted(f1) {
		fmt.Println("The float64 array is sorted")
	} else {
		fmt.Println("The float64 array is NOT sorted")
	}
}
