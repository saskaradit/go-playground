package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `jengjet`
	byteSlice, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}

	login := `jengjsset`
	fmt.Println(s)
	fmt.Println(byteSlice)

	err = bcrypt.CompareHashAndPassword(byteSlice, []byte(login))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("you are logged in")

}
