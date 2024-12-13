package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Book struct {
	title    string
	price    float64
	quantity int
}

func mainReadCsv() {
	bks := make([]Book, 1)
	file, err := os.Open("./chapter_12/products.txt")
	if err != nil {
		log.Fatalf("Error %s opening file products.txt: ", err.Error())
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		// read one line from the file
		line, err := reader.ReadString('\n')
		readErr := err
		// remove the \n from the line
		line = line[:len(line)-1]

		strs := strings.Split(line, ";")
		book := new(Book)
		book.title = strs[0]
		book.price, err = strconv.ParseFloat(strs[1], 64)
		if err != nil {
			fmt.Printf("Error in file: %v", err)
		}
		book.quantity, err = strconv.Atoi(strs[2])
		if err != nil {
			fmt.Printf("Error in file: %v", err)
		}

		if bks[0].title == "" {
			bks[0] = *book
		} else {
			bks = append(bks, *book)
		}

		if readErr == io.EOF {
			break
		}
	}
	fmt.Println("We have read the following books from the file: ")
	for _, bk := range bks {
		fmt.Println(bk)
	}
}
