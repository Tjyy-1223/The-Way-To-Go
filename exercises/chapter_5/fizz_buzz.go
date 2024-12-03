package main

import "fmt"

func mainFizzBuzz() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%3 == 0:
			fmt.Printf("the number %d , print %s\n", i, "Fizz")
		case i%5 == 0:
			fmt.Printf("the number %d , print %s\n", i, "Buzz")
		case i%3 == 0 && i%5 == 0:
			fmt.Printf("the number %d , print %s\n", i, "FizzBuzz")
		default:
			fmt.Println(i)
		}
	}
}
