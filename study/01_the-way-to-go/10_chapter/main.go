package main

import (
	"fmt"
	"go_code/study/01_the-way-to-go/struct_pack"
)

func main() {
	struct1 := new(struct_pack.ExpStruct)
	struct1.Mi1 = 10
	struct1.Mi2 = 16.0

	fmt.Printf("Mi1 = %d\n", struct1.Mi1)
	fmt.Printf("Mi2 = %f\n", struct1.Mi2)
}
