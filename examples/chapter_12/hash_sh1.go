package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
)

func mainSh1() {
	hasher := sha1.New()
	// write string into the 'hasher'
	io.WriteString(hasher, "test")

	var b []byte
	fmt.Printf("Result: %x\n", hasher.Sum(b))
	fmt.Printf("Result: %d\n", hasher.Sum(b))

	hasher.Reset()
	data := []byte("We shall overcome!")
	n, err := hasher.Write(data)
	if n != len(data) || err != nil {
		log.Printf("Hash write error: %v / %v", n, err)
	}

	checksum := hasher.Sum(b)
	fmt.Printf("Result: %x\n", checksum)
}
