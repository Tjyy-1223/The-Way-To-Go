package main

import (
	"container/list"
	"fmt"
)

func mainDlinkedList() {
	lst := list.New()
	lst.PushBack(100)
	lst.PushBack(102)
	lst.PushBack(103)

	for e := lst.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
