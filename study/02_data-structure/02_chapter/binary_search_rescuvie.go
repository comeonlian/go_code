package main

import "fmt"

func binarySearchRescuvie(array []int, target, left, right int) (res int) {
	if left > right {
		return -1
	}
	mid := (left + right) / 2
	midVal := array[mid]

	if target == midVal {
		res = mid
	} else if target < midVal {
		res = binarySearchRescuvie(array, target, left, mid-1)
	} else {
		res = binarySearchRescuvie(array, target, mid+1, right)
	}
	return
}

func main() {
	array := []int{1, 5, 9, 15, 81, 89, 123, 189, 333}
	target := 189
	result := binarySearchRescuvie(array, target, 0, len(array)-1)
	fmt.Println(result)

	target = 8
	result = binarySearchRescuvie(array, target, 0, len(array)-1)
	fmt.Println(result)
}
