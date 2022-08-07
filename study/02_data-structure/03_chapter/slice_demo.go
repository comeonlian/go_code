package main

import "fmt"

func main() {
	array := make([]int, 0, 2)
	fmt.Println("cap: ", cap(array), ", len: ", len(array), ", array: ", array)

	_ = append(array, 1)
	fmt.Println("cap: ", cap(array), ", len: ", len(array), ", array: ", array)
	_ = append(array, 1)
	fmt.Println("cap: ", cap(array), ", len: ", len(array), ", array: ", array)
	_ = append(array, 1)
	fmt.Println("cap: ", cap(array), ", len: ", len(array), ", array: ", array)

	fmt.Println("-----------")

	array = append(array, 1)
	fmt.Println("cap: ", cap(array), ", len: ", len(array), ", array: ", array)
	array = append(array, 1)
	fmt.Println("cap: ", cap(array), ", len: ", len(array), ", array: ", array)
	array = append(array, 1)
	fmt.Println("cap: ", cap(array), ", len: ", len(array), ", array: ", array)

	array = append(array, 1, 1, 1, 1)
	fmt.Println("cap: ", cap(array), ", len: ", len(array), ", array: ", array)
	array = append(array, 1, 1, 1, 1, 1, 1, 1, 1, 1)
	fmt.Println("cap: ", cap(array), ", len: ", len(array), ", array: ", array)
}
