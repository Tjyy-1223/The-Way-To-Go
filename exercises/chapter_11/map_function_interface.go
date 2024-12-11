package main

import "fmt"

type obj interface{}

func mapFunc(f func(obj) obj, list []obj) []obj {
	res := make([]obj, len(list))
	for i, item := range list {
		res[i] = f(item)
	}
	return res
}

func mainMapFunctionInterface() {
	f := func(i obj) obj {
		switch i.(type) {
		case int:
			return i.(int) * 2
		case string:
			return i.(string) + i.(string)
		}
		return i
	}

	isl := []obj{0, 1, 2, 3, 4, 5}
	res1 := mapFunc(f, isl)
	for _, v := range res1 {
		fmt.Println(v)
	}
	println()

	ssl := []obj{"0", "1", "2", "3", "4", "5"}
	res2 := mapFunc(f, ssl)
	for _, v := range res2 {
		fmt.Println(v)
	}
}
