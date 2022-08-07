package main

import "fmt"

func binarySearchLoop(array []int, target, left, right int) int {
	for {
		if left > right {
			return -1
		}
		mid := (left + right) / 2
		midVal := array[mid]

		if target == midVal {
			return mid
		} else if target < midVal {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
}

func main() {
	array := []int{1, 5, 9, 15, 81, 89, 123, 189, 333}
	target := 123
	result := binarySearchLoop(array, target, 0, len(array)-1)
	fmt.Println(result)

	target = 150
	result = binarySearchLoop(array, target, 0, len(array)-1)
	fmt.Println(result)
}
