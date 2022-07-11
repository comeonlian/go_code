package main

import "fmt"

func main() {
	ch := make(chan string)

	go sendData1(ch)
	getData1(ch)
}

func sendData1(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokyo"
	close(ch)
}

func getData1(ch chan string) {
	for {
		input, open := <- ch
		if !open {
			break
		}
		fmt.Println(input)
	}
}