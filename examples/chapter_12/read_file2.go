package main

import (
	"fmt"
	"os"
)

func mainReadFile2() {
	inputFileName := "./chapter_12/products2.txt"
	file, err := os.Open(inputFileName)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	var col1, col2, col3 []string
	for {
		var v1, v2, v3 string
		_, err := fmt.Fscanln(file, &v1, &v2, &v3)
		if err != nil {
			break
		}
		col1 = append(col1, v1)
		col2 = append(col2, v2)
		col3 = append(col3, v3)
	}
	fmt.Println(col1)
	fmt.Println(col2)
	fmt.Println(col3)
}
