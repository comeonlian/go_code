package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var orig string = "1023"
	var newS string

	var an int
	var err error

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)

	// an, err := strconv.Atoi(orig)
	an, err = strconv.Atoi(orig)
	if err != nil {
		fmt.Printf("orig %s is not an integer - exiting with error\n", orig)
		os.Exit(1)
	}

	fmt.Printf("The integer is %d\n", an)

	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)
}
