package main

import (
	"fmt"
	"go_code/study/02_data-structure/structure"
)

func main() {
	node1 := new(structure.LinkNode)
	node1.Data = 1

	node2 := new(structure.LinkNode)
	node2.Data = 2

	node3 := new(structure.LinkNode)
	node3.Data = 3

	node4 := new(structure.LinkNode)
	node4.Data = 4

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4

	curNode := node1
	for {
		if curNode != nil {
			fmt.Printf("%d ", curNode.Data)
			curNode = curNode.Next
			continue
		} else {
			break
		}
	}
}
