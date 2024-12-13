package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func mainMd5() {
	hasher := md5.New()
	io.WriteString(hasher, "test")

	var b []byte
	fmt.Printf("Result: %x\n", hasher.Sum(b))
	fmt.Printf("Result: %d\n", hasher.Sum(b))
}
