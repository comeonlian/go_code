package main

import (
	"fmt"
	"go_code/study/02_data-structure/structure"
)

func main() {
	s := structure.NewSet(10)

	s.Add(3)
	s.Add(7)
	s.Add(2)
	s.Add(6)

	fmt.Printf("list of all items: %v \n", s.List())

	s.Clear()
	fmt.Printf("list is empty: %v \n", s.IsEmpty())

	s.Add(1)
	s.Add(2)
	s.Add(3)
	fmt.Printf("list has 2 item: %v\n", s.Has(2))

	s.Remove(1)
	s.Remove(2)
	fmt.Printf("list of all items: %v\n", s.List())
}
