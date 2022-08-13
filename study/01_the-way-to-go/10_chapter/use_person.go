package main

import (
	"fmt"
	"go_code/study/01_the-way-to-go/person"
)

func main() {
	p := new(person.Person)

	p.SetFirstName("Eric")
	fmt.Println((*p).FirstName())
}
