package main

import "fmt"

func mainSwitch() {
	var num int = 100

	switch num {
	case 98, 99:
		fmt.Println("It's equal to 98,99")
	case 100:
		fmt.Println("It's equal to 100")
	default:
		fmt.Println("It's equal to 98,99 or 100")
	}
}
