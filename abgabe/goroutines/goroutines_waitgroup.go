package main

import (
	"fmt"
	"sync"
)

func talking(from string, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go talking("Hey Alan!", &wg)
	go talking("Hey Philipp!", &wg)

	wg.Wait()
	fmt.Println("done")
}
