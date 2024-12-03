package main

func mainGreeting() {
	println("In main before calling greeting")
	greeting()
	println("In main after calling greeting")
}

func greeting() {
	println("In greeting: Hi!!!")
}
