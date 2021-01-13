package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Username string
	Password string
}

func main() {
	saskarad := user{
		"Saskarad",
		"Jengjet",
	}
	bread := user{
		"bread",
		"Jengjet",
	}

	users := []user{saskarad, bread}
	fmt.Println(users)

	byteSlice, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteSlice))
}
