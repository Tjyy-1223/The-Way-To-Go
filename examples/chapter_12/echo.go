package main

import (
	"flag"
	"fmt"
	"os"
)

var NewLine = flag.Bool("n", false, "print newLine") // echo -n flag, of type *bool

const (
	Space   = " "
	Newline = "\n"
)

func mainEcho() {
	flag.PrintDefaults()
	flag.Parse() // Scans the arg list and sets up flags

	var s string = ""
	fmt.Println(flag.NArg())
	fmt.Println(*NewLine)
	for i := 0; i < flag.NArg(); i++ {
		s += flag.Arg(i)
		s += " "
		if *NewLine { // -n is parsed, flag becomes true
			s += Newline
		}
	}
	os.Stdout.WriteString(s)
}
