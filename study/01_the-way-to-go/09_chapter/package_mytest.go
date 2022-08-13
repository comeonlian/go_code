package main

import (
	"fmt"
	"go_code/study/01_the-way-to-go/mypack"
)

func main() {
	var test1 string
	test1 = mypack.ReturnStr()
	fmt.Printf("ReturnStr from package1: %s \n", test1)
	fmt.Printf("Integer from package1: %d\n", mypack.PackInt)
}
