package main

import "fmt"

func main() {
	ii := []int{1, 2, 3, 4}
	s := sum(ii...)
	fmt.Println(s())

}

func sum(x ...int) func() int {
	total := 0
	return func() int {
		for _, v := range x {
			total += v
		}
		return total
	}
}
