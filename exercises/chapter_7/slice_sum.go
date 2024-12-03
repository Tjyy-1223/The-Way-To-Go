package main

import "fmt"

// Sum 函数接受一个int类型的切片，并返回它们的和。
func Sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func mainSliceSum() {
	// 调用Sum函数并传递一个int类型的切片
	nums := []int{1, 2, 3, 4, 5}
	result := Sum(nums...)
	fmt.Println("The sum is:", result) // 输出：The sum is: 15
}
