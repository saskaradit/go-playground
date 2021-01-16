package main

import "fmt"

// TABLE TEST
func main() {
	fmt.Println("Total: ", myMultiply(1, 2, 3, 4))
	fmt.Println("Total: ", myMultiply(1, 2, 3, 5))
	fmt.Println("Total: ", myMultiply(1, 2, 3, 8))
}

func myMultiply(num ...int) int {
	total := 1
	for v := range num {
		total = total * num[v]
	}
	return total
}
