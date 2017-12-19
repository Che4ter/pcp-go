package main

import (
	"fmt"
	"sort"
)

type sortop func([]int) []int

func Sort(op sortop, unsorted []int) []int {
	return op(unsorted)
}

func ASC(unsorted []int) []int {
	sort.Sort(sort.IntSlice(unsorted))
	return unsorted
}

func DESC(unsorted []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(unsorted)))
	return unsorted
}

func main() {
	unsortedList := []int{1, 5, 3, 7, 2, -3}

	fmt.Println(Sort(ASC, unsortedList))
	fmt.Println(Sort(DESC, unsortedList))
}
