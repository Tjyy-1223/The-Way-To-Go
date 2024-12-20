package main

import "fmt"

type Node struct {
	le   *Node
	data interface{}
	ri   *Node
}

func NewNode(left, right *Node) *Node {
	return &Node{left, nil, right}
}

func (n *Node) SetData(data interface{}) {
	n.data = data
}

func mainNodeStructure() {
	root := NewNode(nil, nil)
	root.SetData("Root Node")

	a := NewNode(nil, nil)
	a.SetData("left node")
	b := NewNode(nil, nil)
	b.SetData("right node")

	root.le = a
	root.ri = b
	fmt.Printf("%v\n", root)
}
