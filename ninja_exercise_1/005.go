package main

import "fmt"

type jengjet int

var x jengjet
var r int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", r)
	x = 42
	r := int(x) + 5
	fmt.Println(x)
	fmt.Println(r)
}
