package main

var a2 = "G"

func mainGlobalScope() {
	n2()
	m2()
	n2()
}

func n2() {
	print(a2)
}

func m2() {
	a2 = "O"
	print(a2)
}
