package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main3till5Char() {
	inputFile, _ := os.Open("./chapter_12/input.dat")
	outputFile, _ := os.OpenFile("./chapter_12/outputT", os.O_WRONLY|os.O_CREATE, 0666)
	defer inputFile.Close()
	defer outputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	outputWriter := bufio.NewWriter(outputFile)
	defer outputWriter.Flush()

	for {
		inputString, _, readerError := inputReader.ReadLine()
		if readerError == io.EOF {
			fmt.Println("EOF")
			return
		}

		fmt.Println(inputString[0:1])
		outputString := string(inputString[0:1]) + "\n"
		_, err := outputWriter.WriteString(outputString)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Println("Conversion done")
}
