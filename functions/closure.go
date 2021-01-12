package main

import "fmt"

func main() {
	ii := []int{1, 2, 3, 4}
	fmt.Println(ii)
	inc := incrementor()
	inc2 := incrementor()

	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc2())
	fmt.Println(inc2())
	{
		closure := 9090
		fmt.Println(closure)
	}
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
