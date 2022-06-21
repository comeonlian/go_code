package main

import "fmt"

type Base struct{}

func (Base) Magic() {
	fmt.Println("base magic")
}

func (self Base) MoreMagic() {
	self.Magic()
	self.Magic()
}

type VooDoo struct {
	Base
}

func (VooDoo) Magic() {
	fmt.Println("voodoo magic")
}

func main() {
	v := new(VooDoo)
	v.Magic()
	v.MoreMagic()
}
