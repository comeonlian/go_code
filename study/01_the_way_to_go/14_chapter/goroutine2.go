package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go sendData(ch)
	go getData(ch)

	time.Sleep(50e9)
}

func sendData(ch chan string) {
	ch <- "Washington"
	time.Sleep(1e9)
	ch <- "Tripoli"
	time.Sleep(1e9)
	ch <- "London"
	time.Sleep(1e9)
	ch <- "Beijing"
	time.Sleep(1e9)
	ch <- "Tokyo"
}

func getData(ch chan string) {
	var input string

	for true {
		input = <-ch
		fmt.Printf("%s ", input)
		time.Sleep(2e9)
	}
}
