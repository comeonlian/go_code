package main

import "fmt"

func rescuvieTail(n, tail int) int {
	if n == 1 {
		return tail
	}
	return rescuvieTail(n-1, n*tail)
}

func main() {
	fmt.Println(rescuvieTail(5, 1))
}
