package main

import (
	"examples/chapter_9/book/pack1"
	"fmt"
)

func mainPackage() {
	var test1 string
	test1 = pack1.ReturnStr()
	fmt.Printf("ReturnStr from package1: %s\n", test1)
	fmt.Printf("Integer from package1: %d\n", pack1.Pack1Int)
}
