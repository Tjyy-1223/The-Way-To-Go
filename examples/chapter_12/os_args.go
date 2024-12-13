package main

import (
	"fmt"
	"os"
	"strings"
)

func mainAlice() {
	who := "Alice "
	if len(os.Args) > 1 {
		who += strings.Join(os.Args[1:], " ")
	}
	fmt.Println("Good Morning", who)
}
