package main

import "fmt"

type person struct {
	firstName       string
	lastName        string
	favoriteFlavour string
}

func main() {
	Rad1 := person{
		firstName:       "Rad",
		lastName:        "Saskara",
		favoriteFlavour: "Vanilla",
	}
	fmt.Println(Rad1)
}
