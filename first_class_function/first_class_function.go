package main

import (
	"fmt"
	"time"
)

func sumUp() func(int) int {
	base := 0 // HL

	return func(value int) int {
		base += value // HL
		<-time.After(time.Millisecond * 1000)
		return base
	}
}

func main() {
	f := sumUp()

	fmt.Println(f(3))
	fmt.Println(f(7))
	fmt.Println(f(14))
}
