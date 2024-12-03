package main

import (
	"fmt"
	"os"
	"runtime"
)

func mainGoos() {
	var goos string = runtime.GOOS
	fmt.Printf("The oprating system is: %s \n", goos)

	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)

	var path2 string = os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path2)
}
