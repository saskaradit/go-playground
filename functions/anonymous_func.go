package main

import "fmt"

func main() {

	func(x int) {
		fmt.Println(x)
	}(45)

	//func expression

	f := func(x int) {
		fmt.Println(x)
	}
	f(42)

	x := func() int {
		return 451
	}()

	y := barz()
	i := y()
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(i)
	// returning a func
}

func foo() string {
	s := "Hello World"
	return s
}

func barz() func() int {
	return func() int {
		return 23
	}
}
