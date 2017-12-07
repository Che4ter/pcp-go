package main

import (
	"fmt"
	"sync"
)

func f(from string, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go f("Hey Alan!", &wg)
	go f("Hey Philipp!", &wg)

	wg.Wait()
	fmt.Println("done")
}
