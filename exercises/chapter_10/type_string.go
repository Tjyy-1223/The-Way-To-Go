package main

import "fmt"

type T struct {
	a int
	b float32
	c string
}

func (t *T) String() string {
	return fmt.Sprintf("%d / %f / %q", t.a, t.b, t.c)
}

func mainTypeString() {
	t := &T{7, -2.35, "abc\tdef"}
	fmt.Printf("%v\n", t)
}
