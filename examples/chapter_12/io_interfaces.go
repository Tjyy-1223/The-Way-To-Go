package main

import (
	"bufio"
	"fmt"
	"os"
)

func mainIoInterfaces() {
	// unbuffered
	fmt.Fprintf(os.Stdout, "%s\n", "hello world! - unbuffered")
	// buffered: os.Stdout implements io.Writer
	buf := bufio.NewWriter(os.Stdout)
	fmt.Fprintf(buf, "%s\n", "hello world! - buffered")
	buf.Flush()
}
