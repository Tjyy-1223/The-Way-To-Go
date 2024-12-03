package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	res := mapFunc(func(i int) int {
		return i * 10
	}, nums)

	fmt.Println(res)
}

func mapFunc(fn func(int) int, nums []int) []int {
	res := make([]int, len(nums))
	for i := range nums {
		res[i] = fn(nums[i])
	}
	return res
}
