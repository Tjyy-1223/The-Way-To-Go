package main

import "fmt"

func mainMultiReturnVal() {
	sum, prod, diff := SumProductDiff(3, 4)
	fmt.Println("Sum:", sum, "| Product:", prod, "| Diff:", diff)
	sum, prod, diff = SumProductDiffN(3, 4)
	fmt.Println("Sum:", sum, "| Product:", prod, "| Diff:", diff)
}

func SumProductDiff(i, j int) (int, int, int) {
	return i + j, i * j, i - j
}

func SumProductDiffN(i, j int) (s int, p int, d int) {
	s, p, d = i+j, i*j, i-j
	return
}
