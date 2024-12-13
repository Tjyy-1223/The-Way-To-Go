package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	strconv "strconv"
)

func cat(r *bufio.Reader) {
	idx := 0
	for {
		buf, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if *PrintLine {
			idx += 1
			buf = []byte(strconv.Itoa(idx) + " " + string(buf))
		}
		fmt.Fprintf(os.Stdout, "%s", buf)
	}
	return
}

var PrintLine = flag.Bool("n", false, "add a line number")

func mainCatNumbered() {
	flag.Parse()
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}

	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s : error reading from %s: %s\n", os.Args[0], flag.Arg(i), err.Error())
			return
		}
		cat(bufio.NewReader(f))
	}
}
