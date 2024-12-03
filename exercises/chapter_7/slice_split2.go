package main

import "fmt"

func mainSliceSplit2() {
	str := "Google"
	fmt.Println(str[len(str)/2:] + str[:len(str)/2])

	str2 := "Google2"
	fmt.Println(str[len(str2)/2:] + str[:len(str2)/2])
}
