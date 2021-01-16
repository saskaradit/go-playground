package main

import (
	"fmt"
	"go-playground/ninja_exercise_12/002/quote"
	"go-playground/ninja_exercise_12/002/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
