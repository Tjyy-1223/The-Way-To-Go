package main

import "fmt"

const c = "C"

var v int = 5

type T struct{}

func init() { // initializing of package

}

func mainGoTemplate() {
	var a int
	Func1()
	// ...
	fmt.Println(a)
}

func (t T) Method1() {
	// ...
}

func Func1() { // exported function Func1
	// ...
}
