package main

import "fmt"

type person struct {
	name string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("I am", p.name)
}

func saySomething(h human) {
	h.speak()
}

func main() {
	rad := person{"Rad"}
	saySomething(&rad)
}
