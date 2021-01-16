package main

import "testing"

func TestMyMultiply(t *testing.T) {
	if myMultiply(1, 2, 3) != 6 {
		t.Error("Expected 6 got", myMultiply(1, 2, 3))
	}
}
