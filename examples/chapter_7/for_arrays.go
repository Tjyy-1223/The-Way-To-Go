package main

import "fmt"

func mainForArrays() {
	var arr [5]int

	for i := 0; i < len(arr); i++ {
		arr[i] = i * 2
	}

	for i := 0; i < len(arr); i++ {
		fmt.Printf("Array ar index %d is %d\n", i, arr[i])
	}

	for i, i2 := range arr {
		fmt.Printf("Array ar index %d is %d\n", i, i2)
	}

	a := [...]string{"a", "b", "c", "d"}
	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}

}
