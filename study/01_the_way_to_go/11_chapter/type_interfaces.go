package main

import (
	"fmt"
	"math"
)

type Square2 struct {
	side float32
}

type Circle2 struct {
	radius float32
}

type Shaper2 interface {
	Area2() float32
}

func main() {
	var shaper Shaper2

	sq1 := new(Square2)
	sq1.side = 5

	shaper = sq1

	if t, ok := shaper.(*Square2); ok {
		fmt.Printf("The type of shaper is: %T\n", t)
	}

	if u, ok := shaper.(*Circle2); ok {
		fmt.Printf("The type of shaper is: %T\n", u)
	} else {
		fmt.Println("shaper does not contain a variable of type Circle")
	}
}

func (s Square2) Area2() float32 {
	return s.side * s.side
}

func (c Circle2) Area2() float32 {
	return c.radius * c.radius * math.Pi
}
