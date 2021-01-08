package main

import "fmt"

type person struct {
	firstName       string
	lastName        string
	favoriteFlavour string
}

func main() {
	rad1 := person{
		firstName:       "Rad",
		lastName:        "Saskara",
		favoriteFlavour: "Vanilla",
	}
	rad2 := person{
		firstName:       "Dar",
		lastName:        "Araksas",
		favoriteFlavour: "Vanilla",
	}

	m := map[string]person{
		rad1.lastName: rad1,
		rad2.lastName: rad2,
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

	// fmt.Println(rad1)
	// fmt.Println(rad2)
	fmt.Println(m)
}
