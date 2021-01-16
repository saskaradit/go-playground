package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	type test struct {
		data   []int
		answer float64
	}
	tests := []test{
		test{[]int{1, 2, 3}, 2},
		test{[]int{1, 2, 3}, 2},
		test{[]int{1, 2, 3}, 2},
	}

	for _, v := range tests {
		f := CenteredAvg(v.data)
		if f != v.answer {
			t.Error("Got", f, "expected", v.answer)
		}
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1, 2, 3}))
	// Ouput:
	// 2
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10000})
	}
}
