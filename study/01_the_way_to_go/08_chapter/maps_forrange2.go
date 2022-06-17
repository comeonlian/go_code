package main

import "fmt"

func main() {
	items := make([]map[int]int, 5)
	for i := range items {
		items[i] = make(map[int]int, 2)
		items[i][121] = 232
		items[i][66] = 88
	}
	fmt.Printf("Version A: Value of items: %v\n", items)
}
