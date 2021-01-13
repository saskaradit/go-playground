package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	rad := person{
		"rad", "Saskara", 20,
	}
	zeref := person{
		"zeref", "Saskara", 20,
	}
	people := []person{rad, zeref}

	fmt.Println(people)

	byteSlice, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteSlice)) //turn slices into json
}
