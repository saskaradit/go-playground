package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := []int{4, 8, 12, 4, 90, 66, 9090, 2}
	strings := []string{"P", "rad", "Brit", "jay"}

	fmt.Println(slice)
	sort.Ints(slice)
	fmt.Println(slice)

	fmt.Println(strings)
	sort.Strings(strings)
	fmt.Println(strings)
}
