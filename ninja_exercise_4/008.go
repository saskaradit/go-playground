package main

import "fmt"

func main() {
	m := map[string]string{
		"James": "phone",
		"Rad":   "love",
	}
	// fmt.Println(m)
	fmt.Println(m["Rad"])
	fmt.Println(m["James"])
}
