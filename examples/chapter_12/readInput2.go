package main

import (
	"bufio"
	"fmt"
	"os"
)

var inputReader *bufio.Reader
var inputs string
var err error

func mainReadInput2() {
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input: ")
	inputs, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("The input was: %s\n", inputs)
	}
}
