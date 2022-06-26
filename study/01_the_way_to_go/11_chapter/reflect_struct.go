package main

import (
	"fmt"
	"reflect"
)

type NotknownType struct {
	s1, s2, s3 string
}

func (n NotknownType) String() string {
	return n.s1 + " - " + n.s2 + " - " + n.s3
}

var secret interface{} = NotknownType{"Ada", "Go", "Oberson"}

func main() {
	value := reflect.ValueOf(secret)
	tp := reflect.TypeOf(secret)

	fmt.Println(tp)
	knd := value.Kind()
	fmt.Println(knd)

	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("Field %d: %v \n", i, value.Field(i))
	}

	res := value.Method(0).Call(nil)
	fmt.Println(res)
}
