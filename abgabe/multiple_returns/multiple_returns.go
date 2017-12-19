package main

import (
	"errors"
	"fmt"
)

func login(name string) (string, error) {
	if name == "alan" {
		return "hello alan", nil
	} else {
		return "bad request", errors.New("unkown user")
	}
}

func main() {
	user, err := login("philipp")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
}
