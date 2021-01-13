package main

import "fmt"

func main() {
	jengjet := 20
	fmt.Println(&jengjet) // & gives you the address
	fmt.Printf("%T\n", &jengjet)
}
