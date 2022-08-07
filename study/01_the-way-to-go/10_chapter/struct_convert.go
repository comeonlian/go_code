package main

import "fmt"

type number struct {
	f float32
}

type nr number

func main() {
	a := number{5.0}
	b := nr{6.0}

	var c = number(b)
	fmt.Println(a, b, c)
}
