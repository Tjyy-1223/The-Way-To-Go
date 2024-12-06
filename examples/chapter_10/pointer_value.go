package main

import "fmt"

type B2 struct {
	thing int
}

func (b *B2) change() {
	b.thing = 1
}

func (b B2) change2() {
	b.thing = 1
}

func (b B2) write() string {
	return fmt.Sprint(b)
}

func mainPointerValue() {
	var b1 B2
	b1.change()
	fmt.Println(b1.write())

	b2 := new(B2) // b2是指针
	b2.change()
	fmt.Println(b2.write())

	b3 := B2{0}
	b3.change2() // 不会起到修改作用
	fmt.Println(b3.write())
}
