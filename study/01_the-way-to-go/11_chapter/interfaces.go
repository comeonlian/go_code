package main

import "fmt"

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (s *Square) Area() float32 {
	return s.side * s.side
}

func main() {
	square := new(Square)
	square.side = 5

	var shaper Shaper
	shaper = square

	fmt.Printf("The square has area: %f\n", shaper.Area())
}
