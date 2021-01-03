package main

import "fmt"

func main() {
	x := 20
	switch {
	case false:
		fmt.Println(false)
	case true:
		fmt.Println(x)
	}

}
