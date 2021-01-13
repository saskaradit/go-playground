package main

import "fmt"

func main() {
	defer foo()
	bar()
}

func foo() {
	defer func() {
		fmt.Println("foo")
	}()
	fmt.Println("jengjet")
}
func bar() {
	fmt.Println("bar")
}
