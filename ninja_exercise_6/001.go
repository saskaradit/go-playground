package main

import "fmt"

func main() {
	n := foo(20)
	x, y := bar(20, "Hello")

	fmt.Println(n)
	fmt.Println(x, y)

}

func foo(z int) int {
	return z + 15
}

func bar(a int, s string) (int, string) {
	return a, s
}
