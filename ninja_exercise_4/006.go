package main

import "fmt"

func main() {
	y := make([]string, 50, 50)
	y = []string{`jengjet`, `jeng`}

	for i, v := range y {
		fmt.Println(i, v)
	}
}
