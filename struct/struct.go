package main

import "fmt"

type friend struct {
	first string
	last  string
}

type closeFriend struct {
	friend
	kind bool
}

func main() {
	friend1 := friend{
		first: "Rad",
		last:  "Bread",
	}

	cf1 := closeFriend{
		friend: friend{
			first: "Jason",
			last:  "No",
		},
		kind: true,
	}
	fmt.Println("jengjet")
	fmt.Println(friend1)
	fmt.Println(cf1)
}
