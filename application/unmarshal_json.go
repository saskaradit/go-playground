//[{"FirstName":"rad","LastName":"Saskara","Age":20},{"FirstName":"zeref","LastName":"Saskara","Age":20}]

package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	s := `[{"FirstName":"rad","LastName":"Saskara","Age":20},{"FirstName":"zeref","LastName":"Saskara","Age":20}]`
	byteSlice := []byte(s)

	var people []person

	err := json.Unmarshal(byteSlice, &people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Print all the data", people)
}
