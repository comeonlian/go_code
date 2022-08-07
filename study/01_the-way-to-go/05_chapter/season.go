package main

import "fmt"

func main() {
	var month int = 11

	switch month {
	case 2, 3, 4:
		fmt.Println("Spring")
	case 5, 6, 7:
		fmt.Println("Summer")
	case 8, 9, 10:
		fmt.Println("Autumn")
	case 11, 12, 1:
		fmt.Println("Winter")
	}


}
