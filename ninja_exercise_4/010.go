package main

import "fmt"

func main() {
	m := map[string][]string{
		"James": []string{"phone"},
		"Rad":   []string{"love"},
	}

	m[`salisbury`] = []string{`steak`, `book`}

	// fmt.Println(m)
	fmt.Println(m["Rad"])
	fmt.Println(m["James"])

	delete(m, "jengjet")
	delete(m, "Rad")
	fmt.Println(m)
}
