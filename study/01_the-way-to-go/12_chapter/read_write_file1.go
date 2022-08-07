package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	inputFileName := "products.txt"
	outputFileName := "products_copy.txt"

	buf, err := ioutil.ReadFile(inputFileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File error: %s\n", err)
		panic(err.Error())
	}

	fmt.Printf("%s\n", string(buf))

	err = ioutil.WriteFile(outputFileName, buf, 0644)
	if err != nil {
		panic(err.Error())
	}
}
