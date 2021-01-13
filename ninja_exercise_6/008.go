package main

import "fmt"

func main() {

	y := whoa()
	fmt.Println(y())
}

func whoa() func() int {
	return func() int {
		return 99
	}
}
