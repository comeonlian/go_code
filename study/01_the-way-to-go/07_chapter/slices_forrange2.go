package main

import "fmt"

func main() {
	seasons := []string{"Spring", "Summer", "Autumn", "Winter"}

	for ix, value := range seasons {
		fmt.Printf("Season %d is: %s\n", ix, value)
	}

	var season string
	for _, season = range seasons {
		fmt.Printf("%s\n", season)
	}
}
