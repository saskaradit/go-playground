package main

import "fmt"

func main() {
	y := 0
	fmt.Println("address y", &y)
	fmt.Println(y)
	foo(&y)
	fmt.Println("y after", y)
}

func foo(x *int) {
	fmt.Println("x before", x)
	fmt.Println(*x)
	*x = 43
	fmt.Println("x after", x)
	fmt.Println(*x)
}
