package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func mainReadWriteFile() {
	inputFile := "./chapter_12/products.txt"
	outputFile := "./chapter_12/products_copy.txt"

	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
	}
	fmt.Printf("%s\n", string(buf))

	err = ioutil.WriteFile(outputFile, buf, 0644) // oct, not hex
	if err != nil {
		panic(err.Error())
	}

}
