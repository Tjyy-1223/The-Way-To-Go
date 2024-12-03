package main

import "fmt"

func f(nums [3]int) {
	for i := range nums {
		fmt.Printf("the value on the index %d and the memory postion: %d\n", i, &nums[i])
	}
}

func mainArrayValue() {
	var nums [3]int
	for i := range nums {
		fmt.Printf("the value on the index %d and the memory postion: %d\n", i, &nums[i])
	}

	f(nums)
}
