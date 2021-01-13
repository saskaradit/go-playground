package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByLastName []user

func (bn ByLastName) Len() int           { return len(bn) }
func (bn ByLastName) Swap(i, j int)      { bn[i], bn[j] = bn[j], bn[i] }
func (bn ByLastName) Less(i, j int) bool { return bn[i].Last < bn[j].Last }

// By Age

type ByAge []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	// your code goes here

	fmt.Println("before name sorted:", users)
	sort.Sort(ByLastName(users))
	fmt.Println(users)

	fmt.Println("before age sorted", users)
	sort.Sort(ByAge(users))
	fmt.Println(users)

	for _, v := range users {
		sort.Strings(v.Sayings)
	}

	fmt.Println(users)
}
