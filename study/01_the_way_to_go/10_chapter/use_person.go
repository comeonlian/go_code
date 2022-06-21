package main

import (
	"fmt"
	"go_code/study/01_the_way_to_go/person"
)

func main() {
	p := new(person.Person)

	p.SetFirstName("Eric")
	fmt.Println((*p).FirstName())
}
