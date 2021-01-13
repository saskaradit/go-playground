package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type ByName []Person

func (bn ByName) Len() int           { return len(bn) }
func (bn ByName) Swap(i, j int)      { bn[i], bn[j] = bn[j], bn[i] }
func (bn ByName) Less(i, j int) bool { return bn[i].Name < bn[j].Name }

// By Age

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	rad := Person{"rad", 20}
	nat := Person{"nat", 37}
	jay := Person{"jay", 19}
	brit := Person{"brit", 34}

	people := []Person{rad, nat, jay, brit}
	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)

}
