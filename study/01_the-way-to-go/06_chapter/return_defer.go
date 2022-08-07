package main

import "fmt"

func f1() (ret int) {
	defer func() {
		ret++
		fmt.Println(ret)
	}()
	return 1
}

func main() {
	fmt.Println(f1())
}
