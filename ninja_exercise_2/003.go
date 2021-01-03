package main

import "fmt"

const (
	a        = "jengjet"
	z string = "jengjet"
)

func main() {
	b := 42 <= 42
	c := 42 >= 42
	d := 42 != 42
	e := 42 < 42
	f := 42 > 42
	fmt.Println(a, b, c, d, e, f, z)
}
