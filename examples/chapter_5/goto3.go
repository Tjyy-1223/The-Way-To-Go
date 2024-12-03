package main

import "fmt"

func mainGoto3() {
	i := 0
	for true {
		if i >= 3 {
			break
		}
		fmt.Println("value of i is:", i)
		i++
	}
	fmt.Println("A statement just after for loop")

	for i := 0; i < 7; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Odd:", i)
	}
}
