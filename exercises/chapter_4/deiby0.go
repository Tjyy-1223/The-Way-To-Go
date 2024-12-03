package main

func mainDivby0() {
	a, b := 10, 0
	c := a / b // panic: runtime error: integer divide by zero

	print(c)
}
