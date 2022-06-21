package main

import "fmt"

type Log1 struct {
	msg string
}

type Custmer1 struct {
	Name string
	Log1
}

func (l *Log1) Add(s string) {
	l.msg += "\n" + s
}

func (l *Log1) String() string {
	return l.msg
}

func (c *Custmer1) String() string {
	return c.Name + "\nLog: " + fmt.Sprintln(c.Log1.String())
}

func main() {
	c := &Custmer1{"Barak Obama", Log1{"1 - Yes we can!"}}
	c.Add("2 - After me the world will be a better place!")
	fmt.Println(c)
}
