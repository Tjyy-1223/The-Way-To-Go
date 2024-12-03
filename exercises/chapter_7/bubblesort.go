package main

import "fmt"

func mainBubbleSort() {
	nums := []int{2, 6, 4, -10, 8, 89, 12, 68, -45, 37}
	fmt.Println("before sort: ", nums)

	bubblesort(nums)
	fmt.Println("after sort: ", nums)
}

func bubblesort(nums []int) {
	for pass := 1; pass < len(nums); pass++ {
		for i := 0; i < len(nums)-pass; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
	}
}
