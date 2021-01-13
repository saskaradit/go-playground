package main

import "fmt"

type person struct {
	name string
}

func main() {
	rad := person{"rad"}
	fmt.Println(rad)
	changeMe(&rad)
	fmt.Println(rad)
}

func changeMe(p *person) {
	p.name = "Bond"
	// (*p).name = "Brit"  --> the same
}
