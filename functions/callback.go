package main

import "fmt"

func main() {
	ii := []int{1, 2, 3, 4}
	s := sum(ii...)
	fmt.Println(s)

	s2 := even(sum, ii...)
	fmt.Println("even Numbers ", s2)

}

func sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func even(f func(x ...int) int, z ...int) int {
	var zi []int
	for _, v := range z {
		if v%2 == 0 {
			zi = append(zi, v)
		}
	}
	return f(zi...)
}
