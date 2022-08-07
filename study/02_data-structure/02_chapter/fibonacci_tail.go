package main

import "fmt"

func fibonacciTail(n, tail1, tail2 int) int {
	if n == 0 {
		return tail1
	}
	return fibonacciTail(n-1, tail2, tail1+tail2)
}

func main() {
	fmt.Println(fibonacciTail(5, 1, 1))
}
