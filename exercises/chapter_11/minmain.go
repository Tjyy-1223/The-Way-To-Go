package main

import "fmt"

func ints() {
	data := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	a := IntArray(data)
	m := Min(a)
	fmt.Printf("The minimum of the array is: %v\n", m)
}

func strings() {
	data := []string{"ddd", "eee", "bbb", "ccc", "aaa"}
	a := StringArray(data)
	m := Min(a)
	fmt.Printf("The minimum of the array is: %v\n", m)
}

func mainMIN() {
	ints()
	strings()
}
