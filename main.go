package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	x, y := 0, 1
	return func() int {
		x, y = y, x+y
		return x
	}
}

func Calcfibonacci(number int) {
	f := fibonacci()
	for i := 0; i < number; i++ {
		fmt.Println(f())
	}
}

func main() {
	Calcfibonacci(9999)

}
