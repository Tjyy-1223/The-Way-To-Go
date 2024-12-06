package main

import "fmt"

type Log2 struct {
	msg string
}

type Customer2 struct {
	Name string
	Log2
}

func (l *Log2) Add(s string) {
	l.msg += "\n" + s
}

func (l *Log2) String() string {
	return l.msg
}

func (c *Customer2) String() string {
	return c.Name + "\nLog:" + fmt.Sprintln(c.Log2)
}

func mainEmbedFunc2() {
	c := &Customer2{"Barak Obama", Log2{"1 - Yes we can!"}}
	c.Add("2 - After me the world will be a better place!")
	fmt.Println(c)
}
