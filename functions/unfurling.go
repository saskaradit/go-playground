package main

import "fmt"

func main() {
	fmt.Println("jengjet")
	// foo("Rad")
	// line := returnString("Bread")
	// fmt.Println(line)
	xi := []int{2, 3, 4, 5, 6}
	woo("rad", xi...)
}

//variadic parameter
func woo(bond string, x ...int) {
	fmt.Println(x)
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println(i, "adding no ", v, " total is ", sum)
	}
	fmt.Println(bond)
}
