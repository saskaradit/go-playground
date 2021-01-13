package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("not humane", p.first, p.age)
}

func main() {
	rad := person{
		first: "rad",
		last:  "Sas",
		age:   20,
	}
	fmt.Println(rad)
	rad.speak()

}
