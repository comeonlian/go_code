package main

import "fmt"

type Shaper3 interface {
	Area() float32
}

type TopologicalGenus3 interface {
	Rank() int
}

type Square3 struct {
	side float32
}

func (s Square3) Area() float32 {
	return s.side * s.side
}

func (s Square3) Rank() int {
	return 1
}

type Rectangle3 struct {
	length, width float32
}

func (r Rectangle3) Area() float32 {
	return r.length * r.width
}

func (r Rectangle3) Rank() int {
	return 2
}

func main() {
	s := &Square3{5}
	r := Rectangle3{5, 3}

	shapes := []Shaper3{s, r}
	fmt.Println("Looping through shapes for area ...")

	for n, _ := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is：", shapes[n].Area())
	}

	topgen := []TopologicalGenus3{s, r}
	fmt.Println("Looping through topgen for rank ...")
	for n, _ := range topgen {
		fmt.Println("Shape details: ", topgen[n])
		fmt.Println("Topological Genus of this shape is : ", topgen[n].Rank())
	}
}
