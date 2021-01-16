package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("Rad")
	if s != "Welcome my dear Rad" {
		t.Error("Got", s, "Expected, Welcome my dear Rad")
	}
}

func ExampleGreet() {
	fmt.Println(Greet("Rad"))
	// Output:
	// Welcome my dear Rad
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("Rad")
	}
}
