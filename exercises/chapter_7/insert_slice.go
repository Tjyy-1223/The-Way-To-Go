package main

import "fmt"

func mainInsertSlice() {
	s := []string{"M", "N", "O", "P", "Q", "R"}
	in := []string{"A", "B", "C"}
	res := InsertStringSlice(s, in, 0) // at the front
	fmt.Println(res)                   // [A B C M N O P Q R]

	fmt.Println()
	res = InsertStringSlice(s, in, 3) // [M N O A B C P Q R]
	fmt.Println(res)
}

func InsertStringSlice(slice, insertion []string, index int) []string {
	result := make([]string, len(slice)+len(insertion))
	at := copy(result, slice[:index])
	fmt.Println("the first step: ", result)
	at += copy(result[index:], insertion)
	fmt.Println("the second step: ", result)
	copy(result[at:], slice[index:])
	fmt.Println("the last step: ", result)
	return result
}
