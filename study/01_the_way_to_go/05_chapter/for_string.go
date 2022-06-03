package main

import "fmt"

func main() {
	str := "Go is a beautiful language!"
	fmt.Printf("The length of str is: %d\n", len(str))

	for i := 0; i < len(str); i++ {
		fmt.Printf("Character on position %d is: %c \n", i, str[i])
	}

	str2 := "日本语"
	fmt.Printf("The length of str2 is: %d\n", len(str2))
	for i := 0; i < len(str2); i++ {
		fmt.Printf("Character on position %d is: %c \n", i, str2[i])
	}
}
