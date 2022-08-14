package main

import (
	"fmt"
	"go_code/study/02_data-structure/structure"
)

func main() {
	root := &structure.TreeNode{Data: "A"}
	root.Left = &structure.TreeNode{Data: "B"}
	root.Right = &structure.TreeNode{Data: "C"}

	root.Left.Left = &structure.TreeNode{Data: "D"}
	root.Left.Right = &structure.TreeNode{Data: "E"}

	root.Right.Right = &structure.TreeNode{Data: "F"}

	fmt.Println("先序遍历: ")
	structure.PreOrder(root)
	fmt.Println()

	fmt.Println("中序遍历: ")
	structure.MidOrder(root)
	fmt.Println()

	fmt.Println("后续遍历: ")
	structure.PostOrder(root)
	fmt.Println()
}
