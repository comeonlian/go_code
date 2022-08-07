package structure

import (
	"fmt"
	"sync"
)

type Array struct {
	array []int
	len   int
	cap   int
	lock  sync.Mutex
}

func Make(len, cap int) *Array {
	arr := new(Array)
	if len > cap {
		panic("len larger than cap")
	}

	array := make([]int, cap, cap)

	arr.array = array
	arr.len = 0
	arr.cap = cap
	return arr
}

func (arr *Array) Append(element int) {
	arr.lock.Lock()
	defer arr.lock.Unlock()

	if arr.len == arr.cap {
		newCap := 2 * arr.cap

		if arr.cap == 0 {
			newCap = 1
		}

		newArray := make([]int, newCap, newCap)
		for index, v := range arr.array {
			newArray[index] = v
		}

		arr.array = newArray
		arr.cap = newCap
	}

	arr.array[arr.len] = element
	arr.len = arr.len + 1
}

func (arr *Array) AppendMany(element ...int) {
	for _, v := range element {
		arr.Append(v)
	}
}

func (arr *Array) Get(index int) int {
	if arr.len == 0 || index >= arr.len {
		panic("index over len")
	}
	return arr.array[index]
}

func (arr *Array) Len() int {
	return arr.len
}

func (arr *Array) Cap() int {
	return arr.cap
}

func Print(arr *Array) (result string) {
	result = "["
	for i := 0; i < arr.Len(); i++ {
		if i == 0 {
			result = fmt.Sprintf("%s%d", result, arr.Get(i))
		} else {
			result = fmt.Sprintf("%s %d", result, arr.Get(i))
		}
	}
	result = result + "]"
	return
}

