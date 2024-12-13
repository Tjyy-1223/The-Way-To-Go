package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func mainFileInput() {
	// 获取当前工作目录
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}
	fmt.Println(dir)

	inputFile, inputError := os.Open("./chapter_12/input.dat")
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		fmt.Printf("The input was : %s", inputString)
		if readerError == io.EOF {
			return
		}
	}
}
