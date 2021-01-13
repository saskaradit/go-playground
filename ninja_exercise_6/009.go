package main

import "fmt"

func main() {

	num := []int{1, 2, 3, 4, 5}
	y := whoa(multiply, num...)
	fmt.Println(y)
}

func multiply(x ...int) int {
	total := 1
	for _, v := range x {
		total *= v
	}
	return total
}

func whoa(funk func(x ...int) int, z ...int) int {
	var sum []int
	total := 0
	for _, v := range z {
		total += v
		sum = append(sum, total)
	}
	return funk(sum...)
}
