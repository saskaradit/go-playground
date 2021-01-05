package main

import "fmt"

func main() {
	m := map[string]int{
		"James": 007,
		"Rad":   205,
	}
	fmt.Println(m)
	fmt.Println(m["Rad"])
	fmt.Println(m["Jengjet"])

	v, ok := m["Jengjet"]
	fmt.Println(v, ok)

	m["todd"] = 99

	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println(m)
	delete(m, "todd")
	fmt.Println(m)
}
