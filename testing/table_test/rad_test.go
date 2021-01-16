package main

import "testing"

func TestMyMultiply(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		test{[]int{4, 3, 5}, 60},
		test{[]int{4, 3, 2}, 24},
		test{[]int{4, 3, 10}, 120},
	}

	for _, v := range tests {
		if myMultiply(v.data...) != v.answer {
			t.Error("Expected", v.answer, "got", myMultiply(1, 2, 3))
		}
	}

}
