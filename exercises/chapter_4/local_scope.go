package main

var a = "G"

func mainLocalScope() {
	n1()
	m1()
	n1()
}

func n1() {
	print(a)
}

func m1() {
	a := "O"
	print(a)
}
