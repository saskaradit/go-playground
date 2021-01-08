package main

import "fmt"

func main() {

	car := struct {
		doors  string
		types  map[string]int
		colors []string
	}{
		doors: "2",
		types: map[string]int{
			"luxury":     200,
			"avantgarde": 250,
			"AMG":        63,
		},
		colors: []string{
			"Black",
			"Grey",
		},
	}
	fmt.Println(car.doors)
	fmt.Println(car.types)
	fmt.Println(car.colors)
}
