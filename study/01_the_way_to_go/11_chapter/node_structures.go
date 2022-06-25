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

func main() {
	root := NewNode(nil, nil)
	root.SetData("root node")

	le := NewNode(nil, nil)
	le.SetData("left Node")

	ri := NewNode(nil, nil)
	ri.SetData("right node")

	root.le = le
	root.ri = ri

	fmt.Printf("%v \n", root)
}
