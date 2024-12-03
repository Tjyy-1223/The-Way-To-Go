package main

import "fmt"

func main() {
	s := "Google"
	fmt.Println("before reverse: ", s)
	s = reverse(s)
	fmt.Println("after reverse: ", s)
}

func reverse(s string) (res string) {
	strs := []byte(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		b := strs[i]
		strs[i] = strs[j]
		strs[j] = b
	}
	res = string(strs)
	return
}
