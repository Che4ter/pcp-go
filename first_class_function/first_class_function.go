package main

import (
	"fmt"
)

func sumUp() func(int) int {
	base := 0

	return func(value int) int {
		base += value
		return base
	}
}

func main() {
	f := sumUp()

	fmt.Println(f(3))
	fmt.Println(f(7))
	fmt.Println(f(14))
}
