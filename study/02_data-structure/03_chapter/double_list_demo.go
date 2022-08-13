package main

import "go_code/study/02_data-structure/structure"

func main() {
	doubleList := new(structure.DoubleList)

	doubleList.AddFromHead("cat")
	doubleList.AddFromHead("dee")
	doubleList.AddFromHead("fgg")

	doubleList.Print()
}
