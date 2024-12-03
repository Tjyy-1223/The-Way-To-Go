package main

var a3 string

func mainFunctionCallsFunction() {
	a3 = "G"
	print(a3)
	f1()
}

func f1() {
	a3 := "O"
	print(a3)
	f2()
}

func f2() {
	print(a3)
}
