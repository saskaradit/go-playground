package main

import "fmt"

func main() {
	fmt.Println("jengjet")
	// foo("Rad")
	// line := returnString("Bread")
	// fmt.Println(line)
	woo(1, 2, 3, 4, 5, 6)
}

func foo(str string) {
	fmt.Println("Foo", str)
}

func returnString(str string) string {
	return fmt.Sprint("heho", str)
}

//variadic parameter
func woo(x ...int) {
	fmt.Println(x)
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println(i, "adding no ", v, " total is ", sum)
	}
}
