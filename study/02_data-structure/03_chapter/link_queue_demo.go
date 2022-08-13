package main

import (
	"fmt"
	"go_code/study/02_data-structure/structure"
)

func main() {
	queue := new(structure.LinkQueue)

	queue.Push("cat")
	queue.Push("dee")
	queue.Push("lili")

	fmt.Println(queue.Pop())

	queue.Push("pipi")
	queue.Push("fifi")

	fmt.Println(queue.Pop())
}
