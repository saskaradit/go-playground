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

func (cf closeFriend) share() {
	fmt.Println("pls", cf.first, cf.last)
}

func (f friend) share() {
	fmt.Println("i am ", f.first, f.last)
}

type human interface {
	share()
}

func bar(h human) {
	fmt.Println("not humane", h)
}

func main() {
	cf1 := closeFriend{
		friend: friend{
			first: "Jason",
			last:  "Junjun",
		},
		kind: true,
	}

	p1 := friend{
		first: "Dr.",
		last:  "Strange",
	}
	fmt.Println(cf1)
	fmt.Println(p1)
	bar(cf1)
	bar(p1)
}
