package main

import "fmt"

func main() {
	// x := type{values} -> composite literal
	// x := []int{12, 22, 53, 34, 15}
	// fmt.Println(x)
	// fmt.Println(len(x))
	// fmt.Println(cap(x))

	// Slicing
	// for i, v := range x {
	// 	fmt.Println(i, v)
	// }

	//Slicing a slice
	// fmt.Println(x[2:3])

	// Append an array
	// x = append(x, 69, 122)
	// y := []int{102, 9090, 333}
	// x = append(x, y...)
	// fmt.Println(x)

	// Delete from array
	// x = append(x[:2], x[4:]...)
	// fmt.Println(x)

	// Make
	// y := make([]int, 10, 100)
	// // make (type, starter length, multiplier)
	// fmt.Println(y, len(y), cap(y))
	// y[0] = 69
	// y[2] = 10
	// fmt.Println(y)

	// multi dimensional array
	jb := []string{"Rad", "Brad", "Bread"}
	mp := []string{"Rayyyyy", "Brayyyyy", "Breayyyyy"}
	fmt.Println(jb)
	fmt.Println(mp)

	md := [][]string{jb, mp}
	fmt.Println(md)

}
