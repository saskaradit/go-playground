package word

import (
	"fmt"
	"go-playground/ninja_exercise_12/002/quote"
	"testing"
)

func TestUseCount(t *testing.T) {
	m := UseCount("one one two")
	for k, v := range m {
		switch k {
		case "one":
			if v != 2 {
				t.Error("got", v, "expected 2")
			}
		case "two":
			if v != 1 {
				t.Error("got", v, "expected 1")
			}
		}
	}
}

func TestCount(t *testing.T) {
	n := Count("jengjet")
	if n != 1 {
		t.Error("got", n, "expected 1")
	}
}

func ExampleCount() {
	fmt.Println(Count("jengjet"))
	// Output:
	//1
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}
func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}

}
