package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var nrchars, nrwords, nrlines int

func mainWordLetterCount() {
	nrchars, nrwords, nrlines = 0, 0, 0
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input, type S to stop")
	for true {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("An error occured: %s\n", err)
		}
		if input == "S\n" {
			fmt.Println("Here are the counts:")
			fmt.Printf("Number of characters: %d\n", nrchars)
			fmt.Printf("Number of words: %d\n", nrwords)
			fmt.Printf("Number of lines: %d\n", nrlines)
			break
		}
		Counters(input)
	}
}

func Counters(s string) {
	nrchars += len(s)
	nrwords += len(strings.Fields(s))
	nrlines++
}
