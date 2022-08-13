package main

import (
	"fmt"
	"go_code/study/02_data-structure/structure"
)

func main() {
	stack := new(structure.ArrayStack)

	stack.Push("cat")
	stack.Push("dog")
	stack.Push("hen")

	fmt.Println("size: ", stack.Size())
	fmt.Println("pop: ", stack.Pop())
	fmt.Println("pop: ", stack.Pop())
	fmt.Println("size: ", stack.Size())

	stack.Push("drag")
	fmt.Println("pop: ", stack.Pop())
}
