package main

import (
	"fmt"
	"unsafe"
)

func mainSizeInt() {
	var i int = 10
	size := unsafe.Sizeof(i)
	fmt.Println("The size of an int is: ", size)
}
