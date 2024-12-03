package main

import "fmt"

func mainForInit() {
	printNums4()
}

func printNums1() {
	for i := 0; i < 5; i++ {
		var v int
		fmt.Printf("%d ", v)
		v = 5
	}
}

func printNums2() {
	for i := 0; ; i++ {
		fmt.Printf("%d ", i)
	}
}

func printNums3() {
	s := ""
	for s != "aaaaa" {
		fmt.Println("Value of s:", s)
		s = s + "a"
	}
}

func printNums4() {
	for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j,
		s = i+1, j+1, s+"a" {
		fmt.Println("Value of i, j, s:", i, j, s)
	}
}
