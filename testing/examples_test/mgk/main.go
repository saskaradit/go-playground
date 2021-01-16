// Package acdc asks if you are ready to rap
package mgk

// Sum adds an unlimited number of value of type int
func Sum(n ...int) int {
	s := 0
	for _, v := range n {
		s += v
	}
	return s
}
