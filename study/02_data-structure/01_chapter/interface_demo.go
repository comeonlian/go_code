package main

import (
	"fmt"
	"reflect"
)

type A interface {
	Println()
}

type B interface {
	Println()
	Printf() int
}

type A1Instance struct {
	Data string
}

func (a1 *A1Instance) Println() {
	fmt.Println("a1: ", a1.Data)
}

type A2Instance struct {
	Data string
}

func (a2 *A2Instance) Println() {
	fmt.Println("a2: ", a2.Data)
}

func (a2 *A2Instance) Printf() int {
	fmt.Println("a2: ", a2.Data)
	return 0
}

func main() {
	var a A

	a = &A1Instance{Data: "i lao you"} // 将具体的结构体赋予该变量
	a.Println()	// 调用接口的方法

	// 断言类型
	if v, ok := a.(*A1Instance); ok {
		fmt.Println(v)
	} else {
		fmt.Println("not a A1")
	}
	fmt.Println(reflect.TypeOf(a).String())

	a = &A2Instance{Data: "i kao you"} // 将具体的结构体赋予该变量
	a.Println() // 调用接口的方法

	// 断言类型
	if v, ok := a.(*A1Instance); ok {
		fmt.Println(v)
	} else {
		fmt.Println("not a A1")
	}
	fmt.Println(reflect.TypeOf(a).String())

	// 定义一个B接口类型的变量
	var b B
	b = &A2Instance{Data: "i shi you"}
	fmt.Println(b.Printf())
}


