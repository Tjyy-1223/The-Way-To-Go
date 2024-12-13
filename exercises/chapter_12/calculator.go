package main

import (
	"bufio"
	"exercises/chapter_12/stack"
	"fmt"
	"os"
	"strconv"
)

func mainCalculator() {
	buf := bufio.NewReader(os.Stdin)
	calc1 := new(stack.Stack)
	fmt.Println("Give a number, an operator (+, -, *, /), or q to stop:")

	for {
		token, err := buf.ReadString('\n')
		if err != nil {
			fmt.Println("input error: ", err)
			return
		}
		token = token[:len(token)-1]
		switch {
		case token == "q":
			fmt.Println("Calculator stopped")
			return
		case "0" <= token && token <= "999999":
			i, _ := strconv.Atoi(token)
			calc1.Push(i)
		case token == "+":
			q := calc1.Pop()
			p := calc1.Pop()
			fmt.Printf("The result of %d %s %d = %d \n", p, token, q, p+q)
		case token == "-":
			q := calc1.Pop()
			p := calc1.Pop()
			fmt.Printf("The result of %d %s %d = %d \n", p, token, q, p-q)
		case token == "*":
			q := calc1.Pop()
			p := calc1.Pop()
			fmt.Printf("The result of %d %s %d = %d \n", p, token, q, p*q)
		case token == "/":
			q := calc1.Pop()
			p := calc1.Pop()
			fmt.Printf("The result of %d %s %d = %d \n", p, token, q, p/q)
		default:
			fmt.Println("No valid input")
		}
	}
}
