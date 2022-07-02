package main

import "os"

func main() {
	os.Stdout.WriteString("hello, world\n")

	f, _ := os.OpenFile("test.txt", os.O_CREATE|os.O_WRONLY, 0666)
	defer f.Close()

	f.WriteString("hello, world in a file \n")
}
