package main

import "fmt"

func main() {

	y := func() int {
		return 55
	}()
	fmt.Println(y)
}
