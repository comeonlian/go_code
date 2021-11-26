package main

import "fmt"

func main() {
	str := "Go is a beautiful language!"
	fmt.Printf("The length of str is: %d\n", len(str))
	for pos, char := range str {
		fmt.Printf("Character on position %d is: %c \n", pos, char)
	}
	fmt.Println()

	str2 := "Chinese: 日本语"
	fmt.Printf("The length of str2 is: %d \n", len(str2))
	for pos, char := range str2 {
		fmt.Printf("character %c starts at byte position %d\n", char, pos)
	}
	fmt.Println()

	fmt.Printf("index int(rune) rune char bytes \n")
	for index, rune := range str2 {
		fmt.Printf("%-2d   %-5d  %U  '%c'  % x \n", index, rune, rune, rune, []byte(string(rune)))
	}
}
