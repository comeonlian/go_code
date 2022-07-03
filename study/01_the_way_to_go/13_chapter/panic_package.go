package main

import (
	"fmt"
	"go_code/study/01_the_way_to_go/parse"
)

func main() {
	var examples = []string{
		"1 2 3 4 5",
		"100 50 25 12.5 6.25",
		"2 + 2 = 4",
		"1st class",
		""}

	for _, example := range examples {
		fmt.Printf("Parsing %q: \n", example)

		nums, err := parse.Parse(example)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(nums)
	}

}
