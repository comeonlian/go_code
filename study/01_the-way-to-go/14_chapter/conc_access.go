package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	Name   string
	salary float64
	chFunc chan func()
}

func NewPerson(name string, salary float64) *Person {
	p := &Person{name, salary, make(chan func())}
	go p.backend()
	return p
}

func (p *Person) backend() {
	for f := range p.chFunc {
		f()
	}
}

// SetSalary :Set salary.
func (p *Person) SetSalary(sal float64) {
	p.chFunc <- func() {
		p.salary = sal
	}
}

// Salary :Retrieve salary.
func (p *Person) Salary() float64 {
	fChan := make(chan float64)
	p.chFunc <- func() {
		fChan <- p.salary
	}
	return <-fChan
}

func (p *Person) String() string {
	return "Person - name is: " + p.Name + " - salary is: " + strconv.FormatFloat(p.Salary(), 'f', 2, 64)
}

func main() {
	bs := NewPerson("Smith Bill", 2500.5)
	fmt.Println(bs)
	bs.SetSalary(4000.66)
	fmt.Println("Salary changed: ")
	fmt.Println(bs)
}