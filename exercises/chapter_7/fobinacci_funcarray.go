package main

import "fmt"

func mainGetFibonacci() {
	res := getFibonacci(5)
	fmt.Println(res)
}

func getFibonacci(n int) (res []int) {
	res = make([]int, n)
	res[0] = 1
	res[1] = 1
	for i := 2; i < n; i++ {
		res[i] = res[i-2] + res[i-1]
	}
	return res
}
