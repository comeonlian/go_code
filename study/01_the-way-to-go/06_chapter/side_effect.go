package main

import "fmt"

func Multiply(a, b int, reply *int) {
	*reply = a * b
}

func main() {
	n := 0
	Multiply(10, 6, &n)
	fmt.Println("Multiply: ", n)
}
