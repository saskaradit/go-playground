package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6}
	a := foo(x...)
	b := foo(x...)

	fmt.Println(a)
	fmt.Println(b)

}

func foo(z ...int) int {
	total := 0
	for _, v := range z {
		total += v
	}
	return total
}

func bar(z ...int) int {
	total := 0
	for _, v := range z {
		total += v
	}
	return total
}
