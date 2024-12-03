package main

import "fmt"

func mainSumAndAverage() {
	// var a = [4]float32 {1.0,2.0,3.0,4.0}
	var a = []float32{1.0, 2.0, 3.0, 4.0}
	fmt.Printf("The sum of the array is: %f\n", Sum2(a))

	var b = []int{1, 2, 3, 4, 5}
	sum, average := SumAndAverage(b)
	fmt.Printf("The sum of the array is: %d, and the average is: %f", sum, average)
}

func Sum2(nums []float32) (res float32) {
	for _, val := range nums {
		res += val
	}
	return
}

func SumAndAverage(nums []int) (sum int, avg float32) {
	for _, val := range nums {
		sum += val
	}
	avg = float32(sum / len(nums))
	return
}
