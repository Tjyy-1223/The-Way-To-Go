package main

import "fmt"

func mainForCharacters() {
	getCharacters()
	getCharacters2()
}

func getCharacters() {
	for i := 0; i < 25; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Printf("G")
		}
		fmt.Println()
	}
}

func getCharacters2() {
	var orig string = "GGGGGGGGGGGGGGGGGGGGGGGGG"
	for i := 0; i < 25; i++ {
		fmt.Println(orig[:i+1])
	}
}
