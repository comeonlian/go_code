package main

import "fmt"

func rescuvie(n int) int {
	if n == 1 {
		return 1
	}
	return n * rescuvie(n-1)
}

func main() {
	fmt.Println(rescuvie(5))
}
