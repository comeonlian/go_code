package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}
	slice1 := arr[1:]
	fmt.Println(arr)
	fmt.Println(slice1)

	slice1 = append(slice1, 5, 6)
	fmt.Println(arr)
	fmt.Println(slice1)
}
