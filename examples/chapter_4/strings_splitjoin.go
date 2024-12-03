package main

import (
	"fmt"
	"strings"
)

func mainStringSplitJoin() {
	str := "The quick brown fox jumps over the lazy dog"
	s1 := strings.Fields(str)
	fmt.Printf("Splited in slice: %v \n", s1)

	for _, val := range s1 {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()

	str2 := "GO1|The ABC of Go|25"
	s2 := strings.Split(str2, "|")
	fmt.Printf("Splitted in slice: %v \n", s2)
	for _, val2 := range s2 {
		fmt.Printf("%s - ", val2)
	}
	fmt.Println()

	str3 := strings.Join(s2, ";")
	fmt.Printf("sl joined by ;: %s \n", str3)
}
