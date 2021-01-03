package main

import "fmt"

func main() {
	a := 69
	fmt.Printf("%v\t%b\t%#x\n", a, a, a)
	b := a << 2
	fmt.Printf("%v\t%b\t%#x\n", b, b, b)
}
