package main

import "fmt"

func mainSliceSplit() {
	str := "Google"
	for i := 0; i <= len(str); i++ {
		a, b := split(str, i)
		fmt.Printf("The string %s split at position %d is: %s / %s\n", str, i, a, b)
	}
}

func split(str string, pos int) (string, string) {
	return str[:pos], str[pos:]
}
