package main

import "fmt"

func main() {
	a := 20
	fmt.Println(&a) // & gives you the address
	fmt.Printf("%T\n", &a)

	b := &a
	fmt.Println(b)
	fmt.Println(*b) // gives you the value of the address
}
