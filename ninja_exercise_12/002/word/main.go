// Package word offers custom functions for strings
package word

import "strings"

// UseCount returns the number of times the word is used
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count return the number of words
func Count(s string) int {
	// write the code for this func
	xs := strings.Fields(s)
	return len(xs)
}
