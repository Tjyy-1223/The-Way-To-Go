package main

import "fmt"

func mainLambdaValue() {
	fv := func() {
		fmt.Println("Hello, world!")
	}
	fv()
	fmt.Printf("The type of fv is %T", fv)
}
