package main

import "fmt"

func mainPrintRec() {
	printrec(10)
}

func printrec(i int) {
	if i < 1 {
		return
	}
	fmt.Printf("%d ", i)
	printrec(i - 1)
}
