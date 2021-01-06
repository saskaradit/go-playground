package main

import "fmt"

func main() {
	x := []int{12, 22, 53, 34, 15}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}
