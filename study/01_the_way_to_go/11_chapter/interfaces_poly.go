package main

import "fmt"

type Shaper1 interface {
	Area() float32
}

type Square1 struct {
	side float32
}

func (s Square1) Area() float32 {
	return s.side * s.side
}

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

func main() {
	r := Rectangle{5, 3}
	q := &Square1{5}

	shapes := []Shaper1{r, q}
	fmt.Println("Loop through shapes for area ...")

	for n, _ := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shap is: ", shapes[n].Area())
	}
}
