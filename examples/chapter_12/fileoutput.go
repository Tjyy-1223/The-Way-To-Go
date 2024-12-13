package main

import (
	"bufio"
	"fmt"
	"os"
)

func mainFileOutput() {
	outputFile, outputError := os.OpenFile("./chapter_12/output.dat", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}
	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)
	outputString := "hello world!\n"
	for i := 0; i < 10; i++ {
		outputWriter.WriteString(outputString)
	}
	err := outputWriter.Flush()
	if err != nil {
		return
	}
}
