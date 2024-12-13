package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

var vc VCard

func mainDegob() {
	file, _ := os.Open("./chapter_12/vcard.gob")
	defer file.Close()

	dec := gob.NewDecoder(file)
	// using the buffer:
	// inReader := bufio.NewReader(file)
	err := dec.Decode(&vc)
	if err != nil {
		log.Println("Error in decoding gob")
	}
	fmt.Println(vc)
}
