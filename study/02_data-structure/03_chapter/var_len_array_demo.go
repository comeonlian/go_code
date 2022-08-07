package main

import (
	"fmt"
	"go_code/study/02_data-structure/structure"
)

func main() {
	arr := structure.Make(0, 2)
	fmt.Println("cap", arr.Cap(), "len", arr.Len(), "array:", structure.Print(arr))

	arr.Append(10)
	fmt.Println("cap", arr.Cap(), "len", arr.Len(), "array:", structure.Print(arr))

	arr.Append(9)
	fmt.Println("cap", arr.Cap(), "len", arr.Len(), "array:", structure.Print(arr))

	arr.AppendMany(8, 7)
	fmt.Println("cap", arr.Cap(), "len", arr.Len(), "array:", structure.Print(arr))
}
