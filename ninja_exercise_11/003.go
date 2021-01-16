package main

import "fmt"

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("there is an error %v", ce.info)
}

func foo(e error) {
	fmt.Println("foo ran - ", e)
}

func main() {
	c1 := customErr{
		info: "need more IQ",
	}
	foo(c1)
}
