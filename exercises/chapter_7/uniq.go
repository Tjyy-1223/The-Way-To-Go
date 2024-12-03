package main

import "fmt"

var arr []byte = []byte{'a', 'b', 'a', 'a', 'a', 'c', 'd', 'e', 'f', 'g'}

func mainUniq() {
	res := make([]byte, len(arr))
	idx := 0
	tmp := byte(0)

	for _, val := range arr {
		if val != tmp {
			res[idx] = val
			fmt.Printf("%c ", res[idx])
			idx++
		}
		tmp = val
	}

}
