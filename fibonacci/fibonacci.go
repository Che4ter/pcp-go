package main

import (
	"fmt"
	"log"
	"math/big"
)

/* Thanks to https://golang.org/ for this fancy function  */
/* But you really should use something bigger than an int */
func fib() func() *big.Int {
	a, b := ToBig(0), ToBig(1)
	
	/* f erhaltet diese Funktion */
	return func() *big.Int {		
		a, b = b, Add(a, b)
		return a
	}
}

/* https://golang.org/pkg/math/big/ */
func ToBig(x int64) *big.Int {
	return big.NewInt(x)
}

/* https://golang.org/pkg/math/big/ */
func Add(x, y *big.Int) *big.Int {
    return big.NewInt(0).Add(x, y)
}

func readInt() int {
	var i int
    _, err := fmt.Scanf("%d", &i)
	
	if err != nil {
		log.Fatal(err)
	}
	
	return i
}

func main() {
	fmt.Printf("Geben Sie eine Zahl ein: ")
	fibN := readInt()
	i := 0
	
	f := fib()
	
	for i < fibN {
		fmt.Print(f(), " ")
		i++
	}
}