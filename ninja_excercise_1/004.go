package main

import "fmt"

type jengjet int

var x jengjet

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
